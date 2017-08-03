package main

import (
	"os"
	"flag"

	"gopkg.in/urfave/cli.v1"
	. "github.com/zenozeng/fossil/proc"
)

func build(c *cli.Context) error {
	//files := make(map[string]string)

	for _, input := range c.Args() {
		FossilDir(Paras{
			input,
			c.String("output"),
			c.Bool("verbose"),
			c.Int("max_io_goroutines"),
			c.String("package")})
	}
	//	const tpl = `package {{.PackageName}}
	//{{range $k, $v := .Files}}
	//const {{$k}} = {{$v}}
	//{{end}}`
	//	t := template.Must(template.New("fossil").Parse(tpl))
	//	t.Execute(os.Stdout, map[string]interface{}{
	//		"PackageName": c.String("package"),
	//		"Files":       files,
	//	})
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
					Name: "bytes",
					Usage: "unimplemented",
				},
				cli.BoolFlag{
					Name:  "verbose",
					Usage: "print verbose information",
				},
				cli.IntFlag{
					Name: "max_io_goroutines, m",
					Usage: "limit max io goroutines",
					Value: 16,
				},
			},
			Action: build,
		},
	}

	app.Run(os.Args)
}
