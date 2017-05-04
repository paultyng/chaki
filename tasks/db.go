package tasks

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

// see https://github.com/jmoiron/sqlx/issues/135
func mapBytesToString(m map[string]interface{}) {
	for k, v := range m {
		if b, ok := v.([]byte); ok {
			m[k] = string(b)
		}
	}
}

func mapScan(rows *sqlx.Rows, dest map[string]interface{}) error {
	err := rows.MapScan(dest)
	if err != nil {
		return err
	}
	mapBytesToString(dest)
	return nil
}

func runSQL(tx *sqlx.Tx, sql string, data map[string]interface{}, maxResults int) (*DBStatementResult, bool, error) {
	maxResultsHit := false
	rx, err := tx.NamedQuery(sql, data)
	if err != nil {
		return nil, maxResultsHit, err
	}
	defer rx.Close()

	sr := &DBStatementResult{
		Data: make([]map[string]interface{}, 0, maxResults),
	}

	rowI := 0
	for rx.Next() {
		rowI++
		if rowI > maxResults {
			maxResultsHit = true
			break
		}

		m := map[string]interface{}{}
		err := mapScan(rx, m)
		if err != nil {
			return nil, maxResultsHit, err
		}

		sr.Data = append(sr.Data, m)
	}

	return sr, maxResultsHit, nil
}

func (t *DBTask) run(data map[string]interface{}, c *Config) (*DBTaskResult, error) {
	dbc, ok := c.DBConnections[t.Connection]
	if !ok {
		return nil, fmt.Errorf("unable to find connection %s", t.Connection)
	}

	conn, err := sqlx.Open(dbc.Driver, dbc.DataSource)
	if err != nil {
		return nil, err
	}

	rollback := true

	tx, err := conn.Beginx()
	if err != nil {
		return nil, err
	}
	defer func() {
		if rollback {
			tx.Rollback()
		}
	}()

	sql := []string(t.SQL)
	res := &DBTaskResult{
		Statements: make([]DBStatementResult, len(sql)),
	}

	for i, s := range sql {
		sr, maxResults, err := runSQL(tx, s, data, c.DBMaxRows)
		if err != nil {
			return nil, err
		}
		if maxResults {
			log.Printf("[WARN] Too many rows returned by statement %d of task", i)
		}
		res.Statements[i] = *sr
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	rollback = false
	return res, nil
}
