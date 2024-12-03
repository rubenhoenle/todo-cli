package cmd

import (
	"fmt"
	"os"

	"git.hoenle.xyz/todo-cli/persistence"
	"github.com/spf13/cobra"
)

func NewRootCommand(dbHandler persistence.DbHandler) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "todo",
		Short: "A simple CLI to manage your TODO items",
		Long:  `A simple CLI to manage your TODO items`,
		Run: func(cmd *cobra.Command, args []string) {
			// show the help to users if no arguments were provided to help them getting started
			if len(args) == 0 {
				cmd.Help()
				return
			}
		},
	}

	rootCmd.AddCommand(newNewCommand(dbHandler))
	rootCmd.AddCommand(newListCommand(dbHandler))
	rootCmd.AddCommand(newShowCommand(dbHandler))
	rootCmd.AddCommand(newResolveCommand(dbHandler))
	rootCmd.AddCommand(newDeleteCommand(dbHandler))

	return rootCmd
}

func Execute(dbHandler persistence.DbHandler) {
	var rootCmd = NewRootCommand(dbHandler)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
