package cmd

import (
	"fmt"
	"os"

	"github.com/infraflakes/srn-cd/pkg/alias"
	"github.com/infraflakes/srn-cd/pkg/tui"
	"github.com/infraflakes/srn-libs/cli"
	"github.com/infraflakes/srn-libs/utils"
	"github.com/spf13/cobra"
)

var RootCmd = cli.NewCommand(
	"scd [alias]",
	"Serein CD - switch directories using aliases and TUI",
	cobra.MaximumNArgs(1),
	func(cmd *cobra.Command, args []string) {
		var target string
		if len(args) == 0 {
			selected, err := tui.RunTUI()
			utils.CheckErr(err)
			if selected == "" {
				os.Exit(0)
			}
			target = selected
		} else {
			target = args[0]
		}

		path, err := alias.Priority(target)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: '%s' is not a valid path or alias\n", target)
			os.Exit(1)
		}

		fmt.Print(path)
	},
)

func Execute() error {
	return RootCmd.Execute()
}
