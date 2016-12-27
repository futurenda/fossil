package main

import (
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "fossil"
	app.Version = "0.0.1"
	app.Usage = "Embedding text file into go constants"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "",
		},
	}
	app.Run(os.Args)
}
