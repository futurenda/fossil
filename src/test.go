package main

import (
	"github.com/zenozeng/fossil/src/proc"
)

func main() {
	//proc.Ls("./input")
	proc.FossilDir(proc.Paras{"./examples/input/", "./examples/output", true})
}
