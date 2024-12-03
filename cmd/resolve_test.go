package cmd

import (
	"bytes"
	"strconv"
	"testing"

	"git.hoenle.xyz/todo-cli/model"
	"git.hoenle.xyz/todo-cli/testutil"

	"github.com/stretchr/testify/assert"
)

func TestResolveWithoutId(t *testing.T) {
	dbHandler := testutil.GetDbHandler("resolve-cmd")
	defer dbHandler.CloseConnection()

	// create the command
	cmd := NewRootCommand(dbHandler)

	// execute the command
	cmd.SetArgs([]string{"resolve"})
	err := cmd.Execute()

	assert.Error(t, err)
	assert.Equal(t, "accepts 1 arg(s), received 0", err.Error())
}

func TestResolveAlreadyResolvedItem(t *testing.T) {
	dbHandler := testutil.GetDbHandler("resolve-cmd")
	defer dbHandler.CloseConnection()

	// create a todo item
	tx, _ := dbHandler.BeginTransaction()
	todoId, _ := tx.CreateTodo(model.Todo{Title: "dummy", Description: "dummy", Resolved: true})
	tx.Commit()

	// create the command
	cmd := NewRootCommand(dbHandler)

	// capture the output
	output := &bytes.Buffer{}
	cmd.SetOut(output)

	// assert the todo item is already resolved beforehand
	todo, _ := dbHandler.GetTodoById(todoId)
	assert.True(t, todo.Resolved)

	// execute the command
	cmd.SetArgs([]string{"resolve", strconv.Itoa(todoId)})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("command execution failed: %v", err)
	}

	assert.Contains(t, output.String(), "TODO #1 is already marked as resolved.")

	// assert the todo item is still marked as resolved afterwards
	todo, _ = dbHandler.GetTodoById(todoId)
	assert.True(t, todo.Resolved)
}

func TestResolve(t *testing.T) {
	dbHandler := testutil.GetDbHandler("resolve-cmd")
	defer dbHandler.CloseConnection()

	// create a todo item
	tx, _ := dbHandler.BeginTransaction()
	todoId, _ := tx.CreateTodo(model.Todo{Title: "dummy", Description: "dummy", Resolved: false})
	tx.Commit()

	// create the command
	cmd := NewRootCommand(dbHandler)

	// capture the output
	output := &bytes.Buffer{}
	cmd.SetOut(output)

	// assert the todo item is unresolved beforehand
	todo, _ := dbHandler.GetTodoById(todoId)
	assert.False(t, todo.Resolved)

	// execute the command
	cmd.SetArgs([]string{"resolve", strconv.Itoa(todoId)})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("command execution failed: %v", err)
	}

	assert.Contains(t, output.String(), "TODO item #1 resolved.")

	// assert the todo item is resolved afterwards
	todo, _ = dbHandler.GetTodoById(todoId)
	assert.True(t, todo.Resolved)
}
