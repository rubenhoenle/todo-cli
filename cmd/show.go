package cmd

import (
	"os"
	"strconv"

	"git.hoenle.xyz/todo-cli/output"
	"git.hoenle.xyz/todo-cli/persistence"
	"github.com/spf13/cobra"
)

func newShowCommand(dbHandler persistence.DbHandler) *cobra.Command {
	var showCmd = &cobra.Command{
		Use:   "show [todo item id]",
		Short: "Show a TODO item",
		Long:  `Show a TODO item with all it's information`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			todoId, err := strconv.Atoi(args[0])
			if err != nil {
				cmd.Printf("Error: '%s' is not a valid todo item id\n", args[0])
				os.Exit(1)
			}

			// TODO: error handling in case the todo item is not found
			todo, err := dbHandler.GetTodoById(todoId)
			if err != nil {
				cmd.Printf("Error loading the todo from the database: %v\n", err)
				os.Exit(1)
			}

			cmd.Print(output.GetPrintStringForTodo(todo))
		},
	}
	return showCmd
}
