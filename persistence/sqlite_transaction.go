package persistence

import (
	"database/sql"
	"git.hoenle.xyz/todo-cli/model"
)

// implements TransactionHandler interface
type SQLiteTransaction struct {
	tx *sql.Tx
}

func (transaction SQLiteTransaction) CreateTodo(todo model.Todo) (int, error) {
	result, err := transaction.tx.Exec("INSERT INTO todo (resolved, title, description) VALUES (?, ?, ?)", todo.Resolved, todo.Title, todo.Description)
	if err != nil {
		return -1, err
	}

	// todo id is provided by the autoincrement of the id column from the database
	todoId, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(todoId), nil
}

func (transaction SQLiteTransaction) UpdateTodo(todo model.Todo) error {
	_, err := transaction.tx.Exec("UPDATE todo SET resolved=?, title=?, description=? WHERE id=?", todo.Resolved, todo.Title, todo.Description, todo.Id)
	return err
}

func (transaction SQLiteTransaction) DeleteTodo(todoId int) error {
	_, err := transaction.tx.Exec("DELETE FROM todo WHERE id=?", todoId)
	return err
}

func (transaction SQLiteTransaction) Commit() error {
	return transaction.tx.Commit()
}

func (transaction SQLiteTransaction) Rollback() error {
	return transaction.tx.Rollback()
}
