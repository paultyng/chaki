package tasks_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/paultyng/chaki/tasks"
)

func TestValidate(t *testing.T) {
	assert := assert.New(t)

	task := &tasks.Task{
		Schema: map[string]interface{}{
			"properties": map[string]interface{}{
				"number": map[string]interface{}{
					"title":   "Order Number",
					"type":    "string",
					"pattern": "[0-9]+",
				},
			},
		},
	}

	cases := []struct {
		number string
		valid  bool
	}{
		{"123", true},
		{"abc", false},
		{"", false},
	}

	for _, c := range cases {
		data := map[string]interface{}{
			"number": c.number,
		}

		err := task.Validate(data)
		if !c.valid {
			_, ok := err.(*tasks.ValidationError)
			assert.True(ok)
			return
		}

		assert.NoError(err)
	}
}
