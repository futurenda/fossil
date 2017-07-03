package main

import (
	sqlFiles "github.com/zenozeng/fossil/examples/output/sql"
	sqlSubFiles "github.com/zenozeng/fossil/examples/output/sql/sub"

	"database/sql"
)

func run(db *sql.DB) {
	db.Exec(sqlFiles.Delete, 1, 2)
	db.Exec(sqlSubFiles.Insert, 1, 2)
}
