package pkg

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/infraflakes/srn-libs/fs"
)

const ConfigFileName = "scd-alias.conf"

// GetConfigPath returns the absolute path to the scd configuration file (~/.config/scd/scd-alias.conf).
func GetConfigPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	scdDir := filepath.Join(configDir, "scd")
	if err := os.MkdirAll(scdDir, 0755); err != nil {
		return "", err
	}
	return filepath.Join(scdDir, ConfigFileName), nil
}

// ReadAliases parses the configuration file and returns a map of alias names to their target paths.
// If the config file does not exist, it returns an empty map and no error.
func ReadAliases() (map[string]string, error) {
	path, err := GetConfigPath()
	if err != nil {
		return nil, err
	}

	aliases := make(map[string]string)
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return aliases, nil
		}
		return nil, err
	}
	defer fs.CloseFile(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			aliases[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}
	return aliases, scanner.Err()
}

// SaveAliases writes the given map of aliases back to the configuration file.
func SaveAliases(aliases map[string]string) error {
	path, err := GetConfigPath()
	if err != nil {
		return err
	}

	file, err := fs.CreateFile(path)
	if err != nil {
		return err
	}
	defer fs.CloseFile(file)

	writer := bufio.NewWriter(file)
	for alias, p := range aliases {
		_, err := fmt.Fprintf(writer, "%s = %s\n", alias, p)
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

// AddAlias associates a name with a specific directory path and saves it to the configuration.
func AddAlias(alias string, path string) error {
	aliases, err := ReadAliases()
	if err != nil {
		return err
	}
	aliases[alias] = path
	return SaveAliases(aliases)
}

// RemoveAlias deletes the alias with the given name from the configuration.
func RemoveAlias(alias string) error {
	aliases, err := ReadAliases()
	if err != nil {
		return err
	}
	if _, ok := aliases[alias]; !ok {
		return fmt.Errorf("alias '%s' not found", alias)
	}
	delete(aliases, alias)
	return SaveAliases(aliases)
}

// WipeAliases clears all saved aliases from the configuration.
func WipeAliases() error {
	return SaveAliases(make(map[string]string))
}

// ExportAliases exports the aliases config to current directory
func ExportAliases(destPath string) error {
	configPath, err := GetConfigPath()
	if err != nil {
		return err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("no aliases available to export")
		}
		return err
	}

	return os.WriteFile(destPath, data, 0755)
}

// FindPathByAlias looks up a target path associated with the given alias name.
func FindPathByAlias(alias string) (string, bool) {
	aliases, err := ReadAliases()
	if err != nil {
		return "", false
	}
	path, ok := aliases[alias]
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
