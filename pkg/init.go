package pkg

import (
	"fmt"
	"os"
)

func GenerateInit(shellName string) (string, error) {
	exe, err := os.Executable()
	if err != nil {
		exe = "scd" // Fallback
	}

	switch shellName {
	case "fish":
		return fmt.Sprintf(`function scd
    set -l target ("%s" $argv)
    if test $status -eq 0; and test -d "$target"
        cd "$target"
    else if test -n "$target"
        printf "%%s\n" $target
    end
end`, exe), nil
	case "bash", "zsh":
		return fmt.Sprintf(`scd() {
    local target
    target=$("%s" "$@")
    if [ $? -eq 0 ] && [ -d "$target" ]; then
        cd "$target"
    else
        [ -n "$target" ] && echo "$target"
    fi
}`, exe), nil
	default:
		return "", fmt.Errorf("unsupported shell: %s. Supported shells: fish, bash, zsh", shellName)
	}
}
