package tasks_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/paultyng/chaki/tasks"
)

func TestNewConfig(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		data string
	}{
		{`{"tasks": { "foo": { "title": "Foo" } } }`},
		{
			`
tasks:
  foo:
    title: Foo
`,
		},
	}

	for i, c := range cases {
		conf, err := tasks.NewConfig(strings.NewReader(c.data))
		assert.NoError(err, "case %d", i)
		assert.NotNil(conf, "case %d", i)
	}
}
