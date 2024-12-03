package cmd

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"git.hoenle.xyz/todo-cli/persistence"

	"github.com/stretchr/testify/assert"
)

func TestHelpOnNoSubcommand(t *testing.T) {
	dbHandler, err := persistence.NewSQLiteHandler("todo-test-db.sqlite3")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer dbHandler.CloseConnection()

	// create the command
	cmd := NewRootCommand(&dbHandler)

	// capture the output
	output := &bytes.Buffer{}
	cmd.SetOut(output)

	// execute the command without arguments
	cmd.SetArgs([]string{})
	err = cmd.Execute()
	if err != nil {
		t.Fatalf("command execution failed: %v", err)
	}

	// validate
	var expectedContains []string
	expectedContains = append(expectedContains, "List all available TODO items")
	expectedContains = append(expectedContains, "Delete a TODO item")
	expectedContains = append(expectedContains, "Create a new TODO item")
	expectedContains = append(expectedContains, "Resolve a TODO item")
	expectedContains = append(expectedContains, "Show a TODO item")
	expectedContains = append(expectedContains, "-h, --help   help for todo")

	assert.True(t, strings.HasPrefix(output.String(), "A simple CLI to manage your TODO items"))
	for _, it := range expectedContains {
		assert.True(t, strings.Contains(output.String(), it))
	}
}
