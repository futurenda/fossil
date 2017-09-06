package main

import (
	"flag"
	"os"

	"github.com/futurenda/fossil/fossil"
	"gopkg.in/urfave/cli.v1"
)

func build(c *cli.Context) error {
	//files := make(map[string]string)

	for _, input := range c.Args() {
		fossil.FossilDir(fossil.FossilParas{
			input,
			c.String("output"),
			c.Bool("verbose"),
			c.Int("max_io_goroutines"),
			c.String("package")})
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
			},
			Action: build,
		},
	}

	app.Run(os.Args)
}
