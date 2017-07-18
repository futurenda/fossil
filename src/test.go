package main

import (
	"github.com/zenozeng/fossil/src/proc"
)

func main() {
	//proc.Ls("./input")
	for _, i := range proc.FossilDir("./examples/input") {
		println(i)
	}
}
