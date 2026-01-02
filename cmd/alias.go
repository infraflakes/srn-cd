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

		if len(aliases) == 0 {
			fmt.Fprintln(os.Stderr, "No aliases configured.")
			return
		}

		for k, v := range aliases {
			fmt.Printf("%s = %s\n", k, v)
		}
	},
)

var aliasExportCmd = cli.NewCommand(
	"export",
	"Export aliases to the current directory",
	cobra.NoArgs,
	func(cmd *cobra.Command, args []string) {
		cwd, err := os.Getwd()
		utils.CheckErr(err)

		dest := filepath.Join(cwd, pkg.ConfigFileName)
		utils.CheckErr(pkg.ExportAliases(dest))
		fmt.Fprintf(os.Stderr, "Exported aliases to: %s\n", dest)
	},
)

var aliasDeleteCmd = cli.NewCommand(
	"delete [name]",
	"Delete an alias",
	cobra.ExactArgs(1),
	func(cmd *cobra.Command, args []string) {
		name := args[0]
		utils.CheckErr(pkg.RemoveAlias(name))
		fmt.Fprintf(os.Stderr, "Deleted alias: %s\n", name)
	},
)

var aliasWipeCmd = cli.NewCommand(
	"wipe",
	"Wipe all aliases",
	cobra.NoArgs,
	func(cmd *cobra.Command, args []string) {
		utils.CheckErr(pkg.WipeAliases())
		fmt.Fprintln(os.Stderr, "All aliases wiped.")
	},
)

func init() {
	aliasCmd.AddCommand(aliasAddCmd)
	aliasCmd.AddCommand(aliasListCmd)
	aliasCmd.AddCommand(aliasDeleteCmd)
	aliasCmd.AddCommand(aliasWipeCmd)
	aliasCmd.AddCommand(aliasExportCmd)
	RootCmd.AddCommand(aliasCmd)
}
