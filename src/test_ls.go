package main

import (
	"github.com/zenozeng/fossil/src/proc"
	"path/filepath"
)

func main() {
	for _, i := range proc.Ls("./examples/") {
		p, _ := filepath.Abs(i.Path)
		println(p)
	}
	for _, i := range proc.Ls("./") {
		p, _ := filepath.Abs(i.Path)
		println(p)
	}
}
