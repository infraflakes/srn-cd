package pkg

import (
	"fmt"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("7")).
			Background(lipgloss.Color("5")).
			Padding(0, 1)

	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("2")).
			Bold(true)

	dimStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("8"))

	columnStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, true, false, false).
			BorderForeground(lipgloss.Color("8")).
			Padding(0, 1)

	currentColumnStyle = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder(), false, true, false, true).
				BorderForeground(lipgloss.Color("5")).
				Padding(0, 1)
)

type entry struct {
	name  string
	isDir bool
	path  string
}

type model struct {
	currentDir     string
	selectedIdx    int
	parentEntries  []entry
	currentEntries []entry
	previewEntries []entry
	width, height  int
	quitting       bool
	finalPath      string
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) updateEntries() {
	m.currentEntries = listEntries(m.currentDir)
	if m.selectedIdx >= len(m.currentEntries) {
		m.selectedIdx = 0
	}
	if len(m.currentEntries) == 0 {
		m.selectedIdx = 0
	}

	m.parentEntries = listEntries(filepath.Dir(m.currentDir))

	if len(m.currentEntries) > 0 {
		sel := m.currentEntries[m.selectedIdx]
		if sel.isDir {
			m.previewEntries = listEntries(sel.path)
		} else {
			m.previewEntries = nil
		}
	} else {
		m.previewEntries = nil
	}
}

func listEntries(path string) []entry {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil
	}

	var entries []entry
	for _, f := range files {
		entries = append(entries, entry{
			name:  f.Name(),
			isDir: f.IsDir(),
			path:  filepath.Join(path, f.Name()),
		})
	}
	return entries
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

		case "up", "k":
			if m.selectedIdx > 0 {
				m.selectedIdx--
				m.updateEntries()
			}

		case "down", "j":
			if m.selectedIdx < len(m.currentEntries)-1 {
				m.selectedIdx++
				m.updateEntries()
			}

		case "left", "h":
			parent := filepath.Dir(m.currentDir)
			if parent != m.currentDir {
				// Find where the current dir was in the parent
				oldDir := filepath.Base(m.currentDir)
				m.currentDir = parent
				m.updateEntries()
				for i, e := range m.currentEntries {
					if e.name == oldDir {
						m.selectedIdx = i
						break
					}
				}
				m.updateEntries()
			}

		case "right", "l":
			if len(m.currentEntries) > 0 {
				sel := m.currentEntries[m.selectedIdx]
				if sel.isDir {
					m.currentDir = sel.path
					m.selectedIdx = 0
					m.updateEntries()
				}
			}

		case "enter":
			if len(m.currentEntries) > 0 {
				m.finalPath = m.currentEntries[m.selectedIdx].path
				return m, tea.Quit
			}
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	return m, nil
}

func (m *model) renderColumn(entries []entry, selectedIdx int, isActive bool) string {
	var s string
	height := m.height - 4 // Leave room for header and footer

	start := 0
	if selectedIdx >= height {
		start = selectedIdx - height + 1
	}

	for i := start; i < len(entries) && i < start+height; i++ {
		e := entries[i]
		name := e.name
		if e.isDir {
			name += "/"
		}

		if i == selectedIdx && isActive {
			s += selectedStyle.Render("> "+name) + "\n"
		} else if i == selectedIdx && !isActive {
			s += dimStyle.Render("  "+name) + "\n"
		} else {
			s += "  " + name + "\n"
		}
	}
	return s
}

func (m *model) View() string {
	if m.quitting {
		return ""
	}

	header := headerStyle.Render(m.currentDir)

	colWidth := (m.width - 6) / 3

	parentCol := columnStyle.Width(colWidth).Height(m.height - 4).Render(m.renderColumn(m.parentEntries, -1, false))
	currentCol := currentColumnStyle.Width(colWidth).Height(m.height - 4).Render(m.renderColumn(m.currentEntries, m.selectedIdx, true))
	previewCol := lipgloss.NewStyle().Width(colWidth).Height(m.height - 4).Render(m.renderColumn(m.previewEntries, -1, false))

	body := lipgloss.JoinHorizontal(lipgloss.Top, parentCol, currentCol, previewCol)

	footer := dimStyle.Render(fmt.Sprintf("\n %d/%d entries | h/j/k/l: navigate | enter: select | q: quit", m.selectedIdx+1, len(m.currentEntries)))

	return lipgloss.JoinVertical(lipgloss.Left, header, body, footer)
}

// RunTUI starts the interactive directory browser.
// It returns the selected path or an empty string if cancelled.
func RunTUI() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	m := &model{
		currentDir: cwd,
	}
	m.updateEntries()

	// Use os.Stderr for the UI so the shell function doesn't capture it.
	// reserves stdout exclusively for the final selected path.
	p := tea.NewProgram(m, tea.WithAltScreen(), tea.WithOutput(os.Stderr))
	finalModel, err := p.Run()
	if err != nil {
		return "", err
	}

	res := finalModel.(*model)
	if res.finalPath != "" {
		return res.finalPath, nil
	}

	return "", nil
}
