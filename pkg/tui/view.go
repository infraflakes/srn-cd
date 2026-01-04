package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// View renders the terminal interface.
func (m *model) View() string {
	if m.quitting {
		return ""
	}

	header := HeaderStyle.Render(m.currentDir)

	colWidth := (m.width - 6) / 3

	// Render the three columns: Parent, Current (active), and Preview
	parentCol := ColumnStyle.Width(colWidth).Height(m.height - 4).Render(m.renderColumn(m.parentEntries, -1, false, colWidth-2))
	currentCol := CurrentColumnStyle.Width(colWidth).Height(m.height - 4).Render(m.renderColumn(m.currentEntries, m.selectedIdx, true, colWidth-2))
	previewCol := Renderer.NewStyle().Width(colWidth).Height(m.height - 4).Render(m.renderColumn(m.previewEntries, -1, false, colWidth-2))

	body := lipgloss.JoinHorizontal(lipgloss.Top, parentCol, currentCol, previewCol)

	filterStatus := " [DIRS]"
	if m.showFiles {
		filterStatus = " [ALL]"
	}
	footer := DimStyle.Render(fmt.Sprintf("\n %d/%d entries%s | .: toggle files | enter: select | q: quit", m.selectedIdx+1, len(m.currentEntries), filterStatus))

	return lipgloss.JoinVertical(lipgloss.Left, header, body, footer)
}

// renderColumn converts a list of entries into a formatted Lip Gloss string for a column.
func (m *model) renderColumn(entries []entry, selectedIdx int, isActive bool, width int) string {
	var s string
	height := m.height - 4

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
			// Full line highlight: no arrow, padded to width
			s += SelectedStyle.Width(width).Render(" "+name) + "\n"
		} else if i == selectedIdx && !isActive {
			s += DimStyle.Render("  "+name) + "\n"
		} else {
			s += "  " + name + "\n"
		}
	}
	return s
}
