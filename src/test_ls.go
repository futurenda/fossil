package main

import (
	"github.com/zenozeng/fossil/src/proc"
	"path/filepath"
)

func main() {
	sqlFileFilter := func(s string) bool {
		return filepath.Ext(s) == ".sql"
	}
	for _, i := range proc.Ls("./examples/", sqlFileFilter, true) {
		p, _ := filepath.Abs(i.Path)
		println(p)
	}
	for _, i := range proc.Ls("./", sqlFileFilter, true) {
		p, _ := filepath.Abs(i.Path)
		println(p)
	}
}
