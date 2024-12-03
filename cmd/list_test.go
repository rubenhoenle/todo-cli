package cmd

import (
	"bytes"
	"strings"
	"testing"

	"git.hoenle.xyz/todo-cli/model"
	"git.hoenle.xyz/todo-cli/testutil"

	"github.com/stretchr/testify/assert"
)

func TestListNoItems(t *testing.T) {
	dbHandler := testutil.GetDbHandler("list-cmd")
	defer dbHandler.CloseConnection()

	// create the command
	cmd := NewRootCommand(dbHandler)

	// capture the output
	output := &bytes.Buffer{}
	cmd.SetOut(output)

	// execute the command
	cmd.SetArgs([]string{"list"})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("command execution failed: %v", err)
	}

	assert.True(t, strings.Contains(output.String(), "No TODOs available."))
}

func TestListMultipleItems(t *testing.T) {
	dbHandler := testutil.GetDbHandler("list-cmd")
	defer dbHandler.CloseConnection()

	// create the command
	cmd := NewRootCommand(dbHandler)

	tx, _ := dbHandler.BeginTransaction()
	tx.CreateTodo(model.Todo{Title: "My first TODO item", Resolved: true})
	tx.CreateTodo(model.Todo{Title: "A TODO title with some Umlauts äöü", Description: "foo", Resolved: false})
	tx.CreateTodo(model.Todo{Title: "Another TODO item", Description: "bar", Resolved: false})
	tx.Commit()

	// capture the output
	output := &bytes.Buffer{}
	cmd.SetOut(output)

	// execute the command
	cmd.SetArgs([]string{"list"})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("command execution failed: %v", err)
	}

	assert.True(t, strings.Contains(output.String(), "TODO #1 - [RESOLVED]"))
	assert.True(t, strings.Contains(output.String(), "Title: My first TODO item"))
	assert.True(t, strings.Contains(output.String(), "Description:"))
	assert.True(t, strings.Contains(output.String(), "TODO #2 - [OPEN]"))
	assert.True(t, strings.Contains(output.String(), "Title: A TODO title with some Umlauts äöü"))
	assert.True(t, strings.Contains(output.String(), "Description: foo"))
	assert.True(t, strings.Contains(output.String(), "TODO #3 - [OPEN]"))
	assert.True(t, strings.Contains(output.String(), "Title: Another TODO item"))
	assert.True(t, strings.Contains(output.String(), "Description: bar"))
}
