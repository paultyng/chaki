package tasks_test

import (
	"testing"

	"github.com/ghodss/yaml"
	"github.com/stretchr/testify/assert"

	"go.ua-ecm.com/chaki/tasks"
)

func TestOptionalStringArrayUnmarshal(t *testing.T) {
	assert := assert.New(t)

	type complexType struct {
		SQL tasks.OptionalStringArray `json:"sql"`
	}

	cases := []struct {
		data     []byte
		expected tasks.OptionalStringArray
	}{
		{[]byte(`yaml`), tasks.OptionalStringArray{"yaml"}},
		{[]byte(`
- yaml 1
- yaml 2
`), tasks.OptionalStringArray{"yaml 1", "yaml 2"}},
		{[]byte(`null`), nil},
		{[]byte(`"single value"`), tasks.OptionalStringArray{"single value"}},
		{[]byte(`["multiple", "values"]`), tasks.OptionalStringArray{"multiple", "values"}},
	}

	for i, c := range cases {
		var actual tasks.OptionalStringArray
		err := yaml.Unmarshal(c.data, &actual)
		if c.expected != nil {
			assert.NotNil(actual)
		} else {
			assert.Nil(actual)
		}
		assert.NoError(err, "case %d", i)
		if err == nil {
			assert.Equal(c.expected, actual, "case %d", i)
		}
	}
}
