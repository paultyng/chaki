package tasks

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert := assert.New(t)

	conf := &Config{
		DBMaxRows: 10,
		DBConnections: map[string]DBConnection{
			"sqlite": DBConnection{
				Driver:     "sqlite3",
				DataSource: ":memory:",
			},
		},
	}

	cases := []struct {
		sql  OptionalStringArray
		data map[string]interface{}
	}{
		{
			OptionalStringArray{"select 1", "select 2", "select 3"},
			nil,
		},
		{
			OptionalStringArray{"select :foo"},
			map[string]interface{}{
				"foo": "bar",
			},
		},
		{
			OptionalStringArray{"select :foo, :bar", "select :bar", "select :foo"},
			map[string]interface{}{
				"foo": "sss",
				"bar": 123,
			},
		},
	}

	for i, c := range cases {
		db := &DBTask{
			Connection: "sqlite",
			SQL:        c.sql,
		}

		// TODO: look at result object
		_, err := db.run(c.data, conf)
		assert.NoError(err, "case %d", i)
	}
}
