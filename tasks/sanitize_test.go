package tasks_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/paultyng/chaki/tasks"
)

func TestSanitize(t *testing.T) {
	assert := assert.New(t)

	dirty := &tasks.Config{
		DBConnections: map[string]tasks.DBConnection{},
		Tasks: map[string]tasks.Task{
			"foo": tasks.Task{
				Title: "Foo",
				DB: &tasks.DBTask{
					Connection: "bar",
				},
			},
		},
	}

	clean := dirty.Sanitize()

	assert.Nil(clean.DBConnections)
	assert.NotNil(clean.Tasks)

	task, ok := clean.Tasks["foo"]
	assert.True(ok)

	assert.Nil(task.DB)
}
