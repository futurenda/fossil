package main

import (
	"os"
	"path/filepath"

	"flag"

	"io/ioutil"

	"gopkg.in/urfave/cli.v1"
	"text/template"
)

func main() {
	flag.Parse()
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

	app.Commands = []cli.Command{
		{
			Name:  "build",
			Usage: "fossil build [FILE]...",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "package",
					Value: "main",
					Usage: "package name",
				},
			},
			Action: func(c *cli.Context) error {
				files := make(map[string]string)
				for _, pattern := range c.Args() {
					matches, err := filepath.Glob(pattern)
					if err != nil {
						return err
					}
					for _, file := range matches {
						fileInfo, err := os.Stat(file)
						if err != nil {
							return err
						}
						fileMode := fileInfo.Mode()
						if fileMode.IsDir() {
							continue
						}
						contents, err := ioutil.ReadFile(file)
						if err != nil {
							return err
						}
						files[file] = string(contents)
					}
				}
				const tpl = `package {{.PackageName}}
{{range $file, $contents := .Files}}
const {{$file}} = ` + "`{{$contents}}`" + `
{{end}}`
				t := template.Must(template.New("fossil").Parse(tpl))
				t.Execute(os.Stdout, map[string]interface{}{
					"PackageName": c.String("package"),
					"Files":       files,
				})
				return nil
			},
		},
	}

	app.Run(os.Args)
}
