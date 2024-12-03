package main

import (
	"fmt"
	"git.hoenle.xyz/todo-cli/cmd"
	"git.hoenle.xyz/todo-cli/persistence"
	"os"
)

func main() {
	dbHandler, err := persistence.NewSQLiteHandler("todo-db.sqlite3")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer dbHandler.CloseConnection()

	cmd.Execute(&dbHandler)
}
