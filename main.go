package main

import (
	"os"
	"flag"

	"gopkg.in/urfave/cli.v1"
	"github.com/zenozeng/fossil/proc"
)

func build(c *cli.Context) error {
	//files := make(map[string]string)

	for _, input := range c.Args() {
		proc.FossilDir(proc.Paras{
			input,
			c.String("output"),
			c.Bool("verbose"),
			16,
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
	app.Version = "0.0.1"
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
				},
				cli.BoolFlag{
					Name:  "verbose",
					Usage: "",
				},
			},
			Action: build,
		},
	}

	app.Run(os.Args)
}
