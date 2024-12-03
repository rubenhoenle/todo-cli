package output

import (
	"strings"
	"testing"

	"git.hoenle.xyz/todo-cli/model"
	"github.com/stretchr/testify/assert"
)

func TestPrintStringForTodo(t *testing.T) {
	title := "Implement more unit tests"
	description := "More unit tests should be implemented to improve the test coverage"

	t.Run("resolved todo item", func(t *testing.T) {
		todo := model.Todo{Id: 27, Title: title, Description: description, Resolved: true}

		result := GetPrintStringForTodo(todo)

		resultLines := strings.Split(result, "\n")

		assert.Equal(t, 4, len(resultLines))
		assert.Equal(t, "TODO #27 - [RESOLVED]", resultLines[0])
		assert.Equal(t, "Title: "+title, resultLines[1])
		assert.Equal(t, "Description: "+description, resultLines[2])
		assert.Equal(t, "", resultLines[3])
	})

	t.Run("open todo item", func(t *testing.T) {
		todo := model.Todo{Id: 27, Title: title, Description: description, Resolved: false}

		result := GetPrintStringForTodo(todo)

		resultLines := strings.Split(result, "\n")

		assert.Equal(t, 4, len(resultLines))
		assert.Equal(t, "TODO #27 - [OPEN]", resultLines[0])
		assert.Equal(t, "Title: "+title, resultLines[1])
		assert.Equal(t, "Description: "+description, resultLines[2])
		assert.Equal(t, "", resultLines[3])
	})
}
