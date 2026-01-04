package alias

import (
	"os"
	"path/filepath"
)

// FindPathByAlias looks up a target path associated with the given alias name.
func FindPathByAlias(target string) (string, bool) {
	aliases, err := ReadAliases()
	if err != nil {
		return "", false
	}
	path, ok := aliases[target]
	return path, ok
}

// Priority resolves the target directory using a two-step approach:
// 1. Checks if the target is an existing directory path (absolute or relative).
// 2. Checks if the target matches a predefined alias in the configuration.
func Priority(target string) (string, error) {
	// Priority 1: Check if it's an existing directory (absolute or relative)
	if info, err := os.Stat(target); err == nil && info.IsDir() {
		return filepath.Abs(target)
	}

	// Priority 2: Alias resolution
	if path, ok := FindPathByAlias(target); ok {
		return path, nil
	}

	return "", os.ErrNotExist
}
