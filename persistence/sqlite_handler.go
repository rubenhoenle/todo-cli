package persistence

import (
	"database/sql"
	_ "embed"
	"fmt"
	_ "github.com/mattn/go-sqlite3"

	"git.hoenle.xyz/todo-cli/model"
)

//go:embed dbinit.sql
var dbInitSqlString string

// implements DbHandler interface
type SQLiteHandler struct {
	db *sql.DB
}

func NewSQLiteHandler(dbFilePath string) (SQLiteHandler, error) {
	db, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		return SQLiteHandler{}, err
	}

	_, err = db.Exec(dbInitSqlString)
	if err != nil {
		return SQLiteHandler{}, err
	}

	return SQLiteHandler{db: db}, nil
}

func (handler *SQLiteHandler) BeginTransaction() (TransactionHandler, error) {
	tx, err := handler.db.Begin()
	if err != nil {
		return SQLiteTransaction{tx: nil}, err
	}
	return SQLiteTransaction{tx: tx}, nil
}

func (handler *SQLiteHandler) GetTodoById(todoId int) (model.Todo, error) {
	row := handler.db.QueryRow("SELECT resolved, title, description FROM todo WHERE id=?", todoId)
	var todo model.Todo
	err := row.Scan(&todo.Resolved, &todo.Title, &todo.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("Error: TODO with ID #%d is not available.\n", todoId)
			return model.Todo{}, err
		}

		fmt.Println(err)
		return model.Todo{}, err
	}
	todo.Id = todoId
	return todo, nil
}

func (handler *SQLiteHandler) GetAllTodos() ([]model.Todo, error) {
	rows, err := handler.db.Query("SELECT id, resolved, title, description FROM todo")
	if err != nil {
		return []model.Todo{}, nil
	}

	var todos []model.Todo
	for rows.Next() {
		var todo model.Todo
		err := rows.Scan(&todo.Id, &todo.Resolved, &todo.Title, &todo.Description)
		if err != nil {
			return []model.Todo{}, nil
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (handler *SQLiteHandler) CloseConnection() error {
	return handler.db.Close()
}
