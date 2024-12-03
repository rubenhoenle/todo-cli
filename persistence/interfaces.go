package persistence

import "git.hoenle.xyz/todo-cli/model"

type DbHandler interface {
	GetTodoById(int) (model.Todo, error)
	GetAllTodos() ([]model.Todo, error)
	BeginTransaction() (TransactionHandler, error)
	CloseConnection() error
}

type TransactionHandler interface {
	Commit() error
	Rollback() error
	CreateTodo(model.Todo) (int, error)
	UpdateTodo(model.Todo) error
	DeleteTodo(int) error
}
