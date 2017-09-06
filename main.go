package main

import (
	"flag"
	"os"

	"github.com/futurenda/fossil/process"
	"gopkg.in/urfave/cli.v1"
)

func build(c *cli.Context) error {
	for _, input := range c.Args() {
		process.FossilDir(process.FossilParas{
			input,
			c.String("output"),
			c.Bool("verbose"),
			c.Int("max_io_goroutines"),
			c.String("package"),
			c.String("extension")})
	}
	return nil
}

func main() {
	flag.Parse()
	app := cli.NewApp()
	app.Name = "fossil"
	app.Version = "1.0.0"
	app.Usage = "Embedding text file into go constants"

	app.Commands = []cli.Command{
		{
			Name:  "build",
			Usage: "fossil build [FILE]...",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "package",
					Value: "",
					Usage: "package name",
				},
				cli.StringFlag{
					Name:  "output, o",
					Value: "",
					Usage: "output dir",
				},
				cli.BoolFlag{
					Name:  "bytes",
					Usage: "unimplemented",
				},
				cli.BoolFlag{
					Name:  "verbose, v",
					Usage: "print verbose information",
				},
				cli.IntFlag{
					Name:  "max_io_goroutines, m",
					Usage: "limit max io goroutines",
					Value: 16,
				},
				cli.StringFlag{
					Name:  "extension, ext",
					Usage: "file type to process",
					Value: "sql",
				},
			},
			Action: build,
		},
	}

	app.Run(os.Args)
}
