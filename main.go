package main

import (
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"go.ua-ecm.com/chaki/cmd"
)

func main() {
	cmd.Execute()
}
