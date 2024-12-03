package testutil

import (
	"fmt"
	"git.hoenle.xyz/todo-cli/persistence"
	"os"
)

func GetDbHandler(id string) persistence.DbHandler {
	sqliteFilePath := fmt.Sprintf("todo-test-db-%s.sqlite3", id)

	err := os.Remove(sqliteFilePath)
	if err != nil && !os.IsNotExist(err) {
		fmt.Printf("Error deleting the existing database: %v\n", err)
		return nil
	}

	dbHandler, err := persistence.NewSQLiteHandler(sqliteFilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return &dbHandler
}
