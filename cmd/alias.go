package cmd

import (
	"fmt"
	"os"

	"github.com/infraflakes/srn-cd/pkg"
	"github.com/infraflakes/srn-libs/cli"
	"github.com/infraflakes/srn-libs/utils"
	"github.com/spf13/cobra"
)

var aliasCmd = cli.NewCommand(
	"alias",
	"Manage directory aliases",
	cobra.NoArgs,
	func(cmd *cobra.Command, args []string) {
		utils.CheckErr(cmd.Help())
	},
)

var aliasAddCmd = cli.NewCommand(
	"add [name]",
	"Add current directory as an alias",
	cobra.ExactArgs(1),
	func(cmd *cobra.Command, args []string) {
		name := args[0]
		cwd, err := os.Getwd()
		utils.CheckErr(err)

		utils.CheckErr(pkg.AddAlias(name, cwd))
		fmt.Fprintf(os.Stderr, "Added alias: %s -> %s\n", name, cwd)
	},
)

var aliasListCmd = cli.NewCommand(
	"list",
	"List all aliases",
	cobra.NoArgs,
	func(cmd *cobra.Command, args []string) {
		aliases, err := pkg.ReadAliases()
		utils.CheckErr(err)

		for k, v := range aliases {
			fmt.Printf("%s = %s\n", k, v)
		}
	},
)

func init() {
	aliasCmd.AddCommand(aliasAddCmd)
	aliasCmd.AddCommand(aliasListCmd)
	RootCmd.AddCommand(aliasCmd)
}
