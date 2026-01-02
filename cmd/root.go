package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/infraflakes/srn-cd/pkg"
	"github.com/infraflakes/srn-libs/cli"
	"github.com/infraflakes/srn-libs/utils"
	"github.com/spf13/cobra"
)

var RootCmd = cli.NewCommand(
	"scd [alias]",
	"Smart CD - switch directories using aliases",
	cobra.MaximumNArgs(1),
	func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			utils.CheckErr(cmd.Help())
			return
		}

		target := args[0]

		// Priority 1: Check if it's an existing directory (absolute or relative)
		if info, err := os.Stat(target); err == nil && info.IsDir() {
			absPath, err := filepath.Abs(target)
			if err == nil {
				fmt.Print(absPath)
				return
			}
		}

		// Priority 2: Alias resolution
		path, ok := pkg.ResolveAlias(target)
		if ok {
			fmt.Print(path)
			return
		}

		fmt.Fprintf(os.Stderr, "Error: '%s' is not a valid path or alias\n", target)
		os.Exit(1)
	},
)

func Execute() error {
	return RootCmd.Execute()
}
