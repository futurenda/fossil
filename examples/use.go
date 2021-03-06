package main

import (
	sqlFiles "github.com/futurenda/fossil/examples/output/sql"
	sqlSubFiles "github.com/futurenda/fossil/examples/output/sql/sub"

	"database/sql"
)

func run(db *sql.DB) {
	db.Exec(sqlFiles.Delete, 1, 2)
	db.Exec(sqlFiles.Mutilline, 1, 2)
	db.Exec(sqlSubFiles.Insert, 1, 2)
}
