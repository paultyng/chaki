package tasks

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert := assert.New(t)

	conf := &Config{
		DBConnections: map[string]DBConnection{
			"sqlite": DBConnection{
				Driver:     "sqlite3",
				DataSource: ":memory:",
			},
		},
	}

	cases := []struct {
		sql  string
		data map[string]interface{}
	}{
		{
			"select 1",
			nil,
		},
		{
			"select :foo",
			map[string]interface{}{
				"foo": "bar",
			},
		},
		{
			"select :foo, :bar",
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

		err := db.run(c.data, conf)
		assert.NoError(err, "case %d", i)
	}
}
