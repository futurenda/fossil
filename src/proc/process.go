package proc

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"fmt"
	"os"
)

func getFileName(s string) string {
	var extension = filepath.Ext(s)
	return s[0:len(s)-len(extension)]
}

func generateContent(info FileInfoWithPath) string {
	content, err := ioutil.ReadFile(info.Path)
	if err != nil {
		panic(err)
	}
	output := "package " + info.Folder +
		"\n\nconst " + strings.Title(getFileName(info.Info.Name())) + " = \"" + strings.TrimSpace(string(content)) + "\"\n"
	return output
}

func generateGoFile(i FileInfoWithPath, paras Paras) {
	content := generateContent(i)
	outputPath := paras.OutputPath
	// todo Windows \
	if outputPath[len(outputPath)-1:] == "/" {
		outputPath = outputPath[0:len(outputPath)-1]
	}
	outputFolder := outputPath + i.RelativePath
	outputPath = outputFolder + "/" + i.Info.Name() + ".go"

	if paras.Verbose {
		fmt.Println("Output to path: " + outputPath)
	}

	_, err := os.Stat(outputFolder)
	if os.IsNotExist(err) {
		if paras.Verbose {
			fmt.Println("Folder: " + outputFolder + " doesn't exist, creating folder.")
		}

		// todo should make sure that the folder creating won't conflict
		err = os.MkdirAll(outputFolder, 0644)
		if err != nil {
			panic(err)
		}
	}

	err = ioutil.WriteFile(outputPath, []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}

func generateAllFile(info []FileInfoWithPath, paras Paras) {
	pool := make(chan bool, paras.Limit)

	done := make(chan int)
	for _, i := range info {
		pool <- true
		go func(info FileInfoWithPath) {
			generateGoFile(info, paras)
			<-pool
			done <- 1
		}(i)
	}
	for i := 0; i < len(info); i++ {
		<-done
	}
}

type Paras struct {
	Dir        string
	OutputPath string
	Verbose    bool
	Limit      int
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
	generateAllFile(info, paras)
}
