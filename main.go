package main

import (
	"os"
	"path/filepath"

	"flag"

	"log"

	"io/ioutil"

	"gopkg.in/urfave/cli.v1"
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
			Action: func(c *cli.Context) error {
				files := make(map[string][]byte, 0)
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
						files[file], err = ioutil.ReadFile(file)
						if err != nil {
							return err
						}
						log.Printf("File: %v, Size: %v bytes", file, fileInfo.Size())
					}
				}
				return nil
			},
		},
	}

	app.Run(os.Args)
}
