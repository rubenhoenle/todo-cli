package cmd

import (
	"git.hoenle.xyz/todo-cli/output"
	"git.hoenle.xyz/todo-cli/persistence"
	"github.com/spf13/cobra"
)

func newListCommand(dbHandler persistence.DbHandler) *cobra.Command {
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all available TODO items",
		Long:  `List all available TODO items to get an overview of all your TODOs`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: feature idea - add filters

			todos, err := dbHandler.GetAllTodos()
			if err != nil {
				cmd.Printf("Error loading todo items: %v\n", err)
				return
			}

			if len(todos) < 1 {
				cmd.Println("No TODOs available.")
				return
			}

			cmd.Print(output.GetPrintStringForTodoSlice(todos))
		},
	}
	return listCmd
}
