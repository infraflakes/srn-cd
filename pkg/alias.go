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

func GetConfigPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	srnDir := filepath.Join(configDir, "serein")
	if err := os.MkdirAll(srnDir, 0755); err != nil {
		return "", err
	}
	return filepath.Join(srnDir, ConfigFileName), nil
}

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
		_, err := writer.WriteString(fmt.Sprintf("%s = %s\n", alias, p))
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

func AddAlias(alias string, path string) error {
	aliases, err := ReadAliases()
	if err != nil {
		return err
	}
	aliases[alias] = path
	return SaveAliases(aliases)
}

func ResolveAlias(alias string) (string, bool) {
	aliases, err := ReadAliases()
	if err != nil {
		return "", false
	}
	p, ok := aliases[alias]
	return p, ok
}
