package main

import (
	"os"
	"path/filepath"

	"flag"

	"io/ioutil"

	"fmt"
	"strconv"
	"strings"
	"text/template"

	"gopkg.in/urfave/cli.v1"
)

func processContents(contents []byte, bytesMode bool) string {
	if !bytesMode {
		c := string(contents)
		if !strings.Contains(c, "`") {
			return fmt.Sprintf("`%s`", c)
		}
		return fmt.Sprintf("%s", strconv.Quote(string(contents)))
	}
	return "TODO"
}

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
				cli.BoolFlag{
					Name: "bytes",
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
						key := strings.ToTitle(strings.Split(file, ".")[0])
						files[key] = processContents(contents, c.Bool("bytes"))
					}
				}
				const tpl = `package {{.PackageName}}
{{range $k, $v := .Files}}
const {{$k}} = {{$v}}
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
