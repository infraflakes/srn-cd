package tui

import (
	"os"
	"path/filepath"
)

// updateEntries refreshes parent, current, and preview entry lists based on the current state.
func (m *model) updateEntries() {
	m.currentEntries = listEntries(m.currentDir, m.showFiles)
	if m.selectedIdx >= len(m.currentEntries) {
		m.selectedIdx = 0
	}
	if len(m.currentEntries) == 0 {
		m.selectedIdx = 0
	}

	m.parentEntries = listEntries(filepath.Dir(m.currentDir), m.showFiles)

	if len(m.currentEntries) > 0 {
		sel := m.currentEntries[m.selectedIdx]
		if sel.isDir {
			m.previewEntries = listEntries(sel.path, m.showFiles)
		} else {
			m.previewEntries = nil
		}
	} else {
		m.previewEntries = nil
	}
}

// listEntries reads a directory and returns a sorted list of entries.
func listEntries(path string, showFiles bool) []entry {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil
	}

	var entries []entry
	for _, f := range files {
		// Filter out files if showFiles is false
		if !showFiles && !f.IsDir() {
			continue
		}
		entries = append(entries, entry{
			name:  f.Name(),
			isDir: f.IsDir(),
			path:  filepath.Join(path, f.Name()),
		})
	}
	return entries
}
