package tui

import (
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
)

// Update() handles terminal events and updates the model state.
func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

		case ".":
			m.showHidden = !m.showHidden
			m.updateEntries()

		case "backspace":
			m.showFiles = !m.showFiles
			m.updateEntries()

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
			m.goUp()

		case "right", "l":
			m.goIn()

		case "enter":
			if len(m.currentEntries) > 0 {
				sel := m.currentEntries[m.selectedIdx]
				// Only allow selecting directories for cd integration
				if sel.isDir {
					m.finalPath = sel.path
					return m, tea.Quit
				}
			}
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	return m, nil
}

// goUp moves the browser into the parent directory.
func (m *model) goUp() {
	parent := filepath.Dir(m.currentDir)
	if parent != m.currentDir {
		oldDir := filepath.Base(m.currentDir)
		m.currentDir = parent
		m.updateEntries()
		// Re-focus the directory we were just in
		for i, e := range m.currentEntries {
			if e.name == oldDir {
				m.selectedIdx = i
				break
			}
		}
		m.updateEntries()
	}
}

// goIn moves the browser into the selected directory.
func (m *model) goIn() {
	if len(m.currentEntries) > 0 {
		sel := m.currentEntries[m.selectedIdx]
		if sel.isDir {
			m.currentDir = sel.path
			m.selectedIdx = 0
			m.updateEntries()
		}
	}
}
