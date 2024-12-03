package cmd

import (
	"bytes"
	"testing"

	"git.hoenle.xyz/todo-cli/testutil"

	"github.com/stretchr/testify/assert"
)

func TestNewNoTitle(t *testing.T) {
	dbHandler := testutil.GetDbHandler("new-cmd")
	defer dbHandler.CloseConnection()

	// create the command
	cmd := NewRootCommand(dbHandler)

	// execute the command
	cmd.SetArgs([]string{"new"})
	err := cmd.Execute()

	assert.Error(t, err)
	assert.Equal(t, "required flag(s) \"title\" not set", err.Error())
}

func TestNewWithoutDescription(t *testing.T) {
	dbHandler := testutil.GetDbHandler("new-cmd")
	defer dbHandler.CloseConnection()

	// create the command
	cmd := NewRootCommand(dbHandler)

	// capture the output
	output := &bytes.Buffer{}
	cmd.SetOut(output)

	// execute the command
	cmd.SetArgs([]string{"new", "--title", "Some test todo item"})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("command execution failed: %v", err)
	}

	assert.Contains(t, output.String(), "TODO item #1 created.")

	todo, _ := dbHandler.GetTodoById(1)
	assert.Equal(t, 1, todo.Id)
	assert.False(t, todo.Resolved)
	assert.Equal(t, "Some test todo item", todo.Title)
	assert.Equal(t, "", todo.Description)
}

func TestNewWithDescription(t *testing.T) {
	dbHandler := testutil.GetDbHandler("new-cmd")
	defer dbHandler.CloseConnection()

	// create the command
	cmd := NewRootCommand(dbHandler)

	// capture the output
	output := &bytes.Buffer{}
	cmd.SetOut(output)

	// execute the command
	cmd.SetArgs([]string{"new", "--title", "Some test todo item with a description", "--description", "This is the description of the todo item"})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("command execution failed: %v", err)
	}

	assert.Contains(t, output.String(), "TODO item #1 created.")

	todo, _ := dbHandler.GetTodoById(1)
	assert.Equal(t, 1, todo.Id)
	assert.False(t, todo.Resolved)
	assert.Equal(t, "Some test todo item with a description", todo.Title)
	assert.Equal(t, "This is the description of the todo item", todo.Description)
}
