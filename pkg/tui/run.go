package tui

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// RunTUI starts the TUI.
// starts with the current working directory
// and configures to output to Stderr to avoid capturing the UI in shell wrappers.
// Returns the selected path or an empty string if the user cancels.
func RunTUI() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	m := &model{
		currentDir: cwd,
	}
	m.updateEntries()

	// Use os.Stderr for the UI to leave Stdout clean for the shell integration.
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
