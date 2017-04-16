package tasks

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func (t *DBTask) run(data map[string]interface{}, c *Config) error {
	dbc, ok := c.DBConnections[t.Connection]
	if !ok {
		return fmt.Errorf("unable to find connection %s", t.Connection)
	}

	conn, err := sqlx.Open(dbc.Driver, dbc.DataSource)
	if err != nil {
		return err
	}

	tx, err := conn.Beginx()
	if err != nil {
		return err
	}

	sql := []string(t.SQL)

	for _, s := range sql {
		_, err = tx.NamedExec(s, data)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
