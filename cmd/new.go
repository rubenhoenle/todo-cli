package cmd

import (
	"os"

	"git.hoenle.xyz/todo-cli/model"
	"git.hoenle.xyz/todo-cli/persistence"
	"github.com/spf13/cobra"
)

func newNewCommand(dbHandler persistence.DbHandler) *cobra.Command {
	var newCmd = &cobra.Command{
		Use:   "new",
		Short: "Create a new TODO item",
		Long:  `Create a new TODO item to keep track of your new task`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: add nice wizard mode

			title, _ := cmd.Flags().GetString("title")
			description, _ := cmd.Flags().GetString("description")

			todo := model.Todo{Title: title, Description: description, Resolved: false}

			tx, err := dbHandler.BeginTransaction()
			if err != nil {
				cmd.Printf("Error starting database transaction: %v\n", err)
				os.Exit(1)
			}

			todoId, err := tx.CreateTodo(todo)
			if err != nil {
				cmd.Printf("Error inserting the todo into the database: %v\n", err)
				tx.Rollback()
				os.Exit(1)
			}

			err = tx.Commit()
			if err != nil {
				cmd.Printf("Error committing database transaction: %v\n", err)
				os.Exit(1)
			}

			cmd.Printf("TODO item #%d created.\n", todoId)
		},
	}
	newCmd.Flags().StringP("title", "t", "", "The title of the new todo item")
	newCmd.Flags().StringP("description", "d", "", "The description of the new todo item")

	newCmd.MarkFlagRequired("title")

	return newCmd
}
