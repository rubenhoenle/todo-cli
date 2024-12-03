package output

import (
	"fmt"
	"strings"

	"git.hoenle.xyz/todo-cli/model"
)

func GetPrintStringForTodo(todo model.Todo) string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "TODO #%d - [%s]\n", todo.Id, getTodoStatusString(todo.Resolved))
	fmt.Fprintf(&builder, "Title: %s\n", todo.Title)
	fmt.Fprintf(&builder, "Description: %s\n", todo.Description)
	return builder.String()
}

func GetPrintStringForTodoSlice(todos []model.Todo) string {
	var builder strings.Builder
	for idx, todo := range todos {
		fmt.Fprintf(&builder, "%s", GetPrintStringForTodo(todo))

		// add seperator between todo items
		if idx+1 < len(todos) {
			fmt.Fprintf(&builder, "\n")
			fmt.Fprintf(&builder, "===================================\n")
			fmt.Fprintf(&builder, "\n")
		}
	}
	return builder.String()
}

func getTodoStatusString(resolved bool) string {
	if resolved {
		return "RESOLVED"
	}
	return "OPEN"
}
