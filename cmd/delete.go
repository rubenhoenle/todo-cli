package cmd

import (
	"os"
	"strconv"

	"git.hoenle.xyz/todo-cli/persistence"
	"github.com/spf13/cobra"
)

func newDeleteCommand(dbHandler persistence.DbHandler) *cobra.Command {
	var deleteCmd = &cobra.Command{
		Use:   "delete [todo item id]",
		Short: "Delete a TODO item",
		Long:  `Delete a TODO item to remove it completely from your list`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			todoId, err := strconv.Atoi(args[0])
			if err != nil {
				cmd.Printf("Error: '%s' is not a valid todo item id\n", args[0])
				os.Exit(1)
			}

			// TODO: error handling in case the todo item is not found
			_, err = dbHandler.GetTodoById(todoId)
			if err != nil {
				cmd.Printf("Error loading the todo from the database: %v\n", err)
				os.Exit(1)
			}

			tx, err := dbHandler.BeginTransaction()
			if err != nil {
				cmd.Printf("Error starting database transaction: %v\n", err)
				os.Exit(1)
			}

			err = tx.DeleteTodo(todoId)
			if err != nil {
				cmd.Printf("Error deleting the todo in the database: %v\n", err)
				tx.Rollback()
				os.Exit(1)
			}

			err = tx.Commit()
			if err != nil {
				cmd.Printf("Error committing database transaction: %v\n", err)
				os.Exit(1)
			}

			cmd.Printf("TODO item #%d deleted.\n", todoId)
		},
	}
	return deleteCmd
}
