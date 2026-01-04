package tui

import (
	"os"

	"github.com/charmbracelet/lipgloss"
)

var (
	// Renderer tied to os.Stderr to allow shell command substitution to capture stdout.
	Renderer = lipgloss.NewRenderer(os.Stderr)

	// HeaderStyle is used for the current directory path at the top.
	HeaderStyle = Renderer.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("7")).
			Background(lipgloss.Color("5")).
			Padding(0, 1)

	// SelectedStyle is used for the currently highlighted item in the active column.
	SelectedStyle = Renderer.NewStyle().
			Foreground(lipgloss.Color("2")).
			Bold(true)

	// DimStyle is used for non-active columns and footer hints.
	DimStyle = Renderer.NewStyle().
			Foreground(lipgloss.Color("8"))

	// ColumnStyle defines the border and padding for parent and preview columns.
	ColumnStyle = Renderer.NewStyle().
			Border(lipgloss.NormalBorder(), false, true, false, false).
			BorderForeground(lipgloss.Color("8")).
			Padding(0, 1)

	// CurrentColumnStyle defines the highlighted border for the active navigation column.
	CurrentColumnStyle = Renderer.NewStyle().
				Border(lipgloss.NormalBorder(), false, true, false, true).
				BorderForeground(lipgloss.Color("5")).
				Padding(0, 1)
)
