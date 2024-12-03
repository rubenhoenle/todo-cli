package cmd

import (
	"os"
	"strconv"

	"git.hoenle.xyz/todo-cli/persistence"
	"github.com/spf13/cobra"
)

func newResolveCommand(dbHandler persistence.DbHandler) *cobra.Command {
	var resolveCmd = &cobra.Command{
		Use:   "resolve [todo item id]",
		Short: "Resolve a TODO item",
		Long:  `Resolve a TODO item to mark it as completed`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			todoId, err := strconv.Atoi(args[0])
			if err != nil {
				cmd.Printf("Error: '%s' is not a valid todo item id\n", args[0])
				os.Exit(1)
			}

			todo, err := dbHandler.GetTodoById(todoId)
			if err != nil {

				cmd.Printf("Error loading the todo from the database: %v\n", err)
				os.Exit(1)
			}

			if todo.Resolved {
				cmd.Printf("TODO #%d is already marked as resolved.\n", todoId)
				return
			}
			todo.Resolved = true

			tx, err := dbHandler.BeginTransaction()
			if err != nil {
				cmd.Printf("Error starting database transaction: %v\n", err)
				os.Exit(1)
			}

			err = tx.UpdateTodo(todo)
			if err != nil {
				cmd.Printf("Error updating the todo in the database: %v\n", err)
				tx.Rollback()
				os.Exit(1)
			}

			err = tx.Commit()
			if err != nil {
				cmd.Printf("Error committing database transaction: %v\n", err)
				os.Exit(1)
			}

			cmd.Printf("TODO item #%d resolved.\n", todoId)
		},
	}
	return resolveCmd
}
