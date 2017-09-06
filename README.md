# Fossil

Convert text files to Go Constants.

See [Design.md](design.md)

## Usage

```bash
tree input

input
└── sql
    ├── delete.sql
    ├── select.sql
    ├── sub
    │   └── insert.sql
    └── update.sql
```

```bash
fossil build input -o output
```

```bash
tree output

output
└── sql
    ├── delete.sql.go
    ├── select.sql.go
    ├── sub
    │   └── insert.sql.go
    └── update.sql.go
```


```go
package main

import (
	sqlFiles "github.com/futurenda/fossil/examples/output/sql"
	sqlSubFiles "github.com/futurenda/fossil/examples/output/sql/sub"

	"database/sql"
)

func run(db *sql.DB) {
	db.Exec(sqlFiles.Delete, 1, 2)
	db.Exec(sqlSubFiles.Insert, 1, 2)
}
```
