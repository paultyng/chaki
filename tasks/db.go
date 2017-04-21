package tasks

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

func (t *DBTask) run(data map[string]interface{}, c *Config) (*DBTaskResult, error) {
	dbc, ok := c.DBConnections[t.Connection]
	if !ok {
		return nil, fmt.Errorf("unable to find connection %s", t.Connection)
	}

	conn, err := sqlx.Open(dbc.Driver, dbc.DataSource)
	if err != nil {
		return nil, err
	}

	tx, err := conn.Beginx()
	if err != nil {
		return nil, err
	}

	sql := []string(t.SQL)
	res := &DBTaskResult{
		Statements: make([]DBStatementResult, len(sql)),
	}

	for i, s := range sql {
		rx, err := tx.NamedQuery(s, data)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		defer rx.Close()

		sr := DBStatementResult{
			Data: make([]map[string]interface{}, 0, c.DBMaxRows),
		}

		rowI := 0
		for rx.Next() {
			rowI++
			if rowI > c.DBMaxRows {
				log.Printf("[WARN] Too many rows returned by statement %d of task", i)
				break
			}

			m := map[string]interface{}{}
			err := rx.MapScan(m)
			if err != nil {
				return nil, err
			}

			sr.Data = append(sr.Data, m)
		}
		res.Statements[i] = sr
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return res, nil
}
