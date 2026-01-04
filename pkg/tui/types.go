package tui

// entry represents a single file system item (file or directory).
type entry struct {
	name  string // Base name of the entry
	isDir bool   // Whether the entry is a directory
	path  string // Absolute path to the entry
}

// the state of the TUI application.
type model struct {
	currentDir     string  // Currently browsed directory
	selectedIdx    int     // Index of the selected entry in currentEntries
	parentEntries  []entry // Cached entries of the parent directory
	currentEntries []entry // Entries in the current directory
	previewEntries []entry // Entries in the selected directory (for preview)
	width, height  int     // Terminal dimensions
	quitting       bool    // Whether the user is exiting the TUI
	finalPath      string  // The path selected to be returned to the shell
	showFiles      bool    // Whether to show regular files (dirs only by default)
	showHidden     bool    // Whether to show hidden files/directories (false by default)
}
