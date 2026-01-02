package cmd

import (
	"fmt"

	"github.com/infraflakes/srn-cd/pkg"
	"github.com/infraflakes/srn-libs/cli"
	"github.com/infraflakes/srn-libs/utils"
	"github.com/spf13/cobra"
)

var initCmd = cli.NewCommand(
	"init [shell]",
	"Generate shell initialization script (fish, bash, zsh)",
	cobra.ExactArgs(1),
	func(cmd *cobra.Command, args []string) {
		shell := args[0]
		script, err := pkg.GenerateInit(shell)
		utils.CheckErr(err)
		fmt.Println(script)
	},
)

func init() {
	RootCmd.AddCommand(initCmd)
}
