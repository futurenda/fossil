package proc

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"fmt"
)

func getFileName(s string) string {
	var extension = filepath.Ext(s)
	return s[0:len(s)-len(extension)]
}

func generateContent(info FileInfoWithPath) string {
	content, _ := ioutil.ReadFile(info.Path)
	output := "package " + info.Folder +
		"\n\nconst " + strings.Title(getFileName(info.Info.Name())) + " = \"" + strings.TrimSpace(string(content)) + "\"\n"
	return output
}

func generateGoFile(info []FileInfoWithPath, paras Paras) {
	for _, i := range info {
		content := generateContent(i)
		outputPath := paras.OutputPath
		// todo Windows \
		if outputPath[len(outputPath)-1:] == "/" {
			outputPath = outputPath[0:len(outputPath) - 1]
		}
		outputPath += i.RelativePath + "/" + i.Info.Name() + ".go"

		if paras.Verbose {
			fmt.Println(outputPath)
		}

		err := ioutil.WriteFile(outputPath, []byte(content), 0744)
		if err != nil {
			panic(err)
		}

	}
}

type Paras struct {
	Dir string
	OutputPath string
	Verbose    bool
}

func FossilDir(paras Paras) {
	// todo flags
	sqlFileFilter := func(s string) bool {
		return filepath.Ext(s) == ".sql"
	}
	info := Ls(paras.Dir, sqlFileFilter, paras.Verbose)
	if paras.Verbose {
		fmt.Printf("Found %d files\n", len(info))

	}

	generateGoFile(info, paras)
}
