package main

import (
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"github.com/paultyng/chaki/cmd"
)

func main() {
	cmd.Execute()
}
