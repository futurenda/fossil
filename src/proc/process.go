package proc

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func processContents(contents []byte, bytesMode bool) string {
	if !bytesMode {
		c := string(contents)
		if !strings.Contains(c, "`") {
			return fmt.Sprintf("`%s`", c)
		}
		return fmt.Sprintf("%s", strconv.Quote(string(contents)))
	}
	// todo
	return "TODO"
}

func getFileName(s string) string {
	var extension = filepath.Ext(s)
	return s[0:len(s)-len(extension)]
}

func snakeToCamelCase(s string) string {
	output := ""
	for i, c := range s {
		if (i == 0 || string(s[i-1]) == "_") && string(s[i]) != "_" {
			output += strings.ToUpper(string(c))
		} else if !(string(s[i]) == "_") {
			output += strings.ToLower(string(c))
		}
	}
	return output
}

func regularizeToSnakeCase(s string) string {
	regularized := ""
	for _, c := range s {
		if unicode.IsLetter(c) {
			if string(c) == strings.ToUpper(string(c)) {
				if len(regularized) != 0{
					regularized += "_"
				}
				regularized += strings.ToLower(string(c))
			} else {
				regularized += strings.ToLower(string(c))
			}
		} else {
			if len(regularized) != 0 || string(c) == "_" {
				regularized += "_"
			}
		}
	}
	return regularized
}

func generateContent(info FileInfoWithPath) string {
	content, err := ioutil.ReadFile(info.Path + info.Name)
	if err != nil {
		panic(err)
	}

	sqlVarName := snakeToCamelCase(regularizeToSnakeCase(getFileName(info.Name)))
	sqlContent := processContents([]byte(strings.TrimSpace(string(content))), false)
	output := fmt.Sprintf("package %s\n\nconst %s = %s\n",
		info.Folder, sqlVarName, sqlContent)
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
	outputPath = outputFolder + "/" + regularizeToSnakeCase(getFileName(i.Name)) + ".sql.go"

	if paras.Verbose {
		fmt.Printf("Output to path: %s\n", outputPath)
	}

	_, err := os.Stat(outputFolder)
	if os.IsNotExist(err) {
		if paras.Verbose {
			fmt.Printf("Folder: %s doesn't exist, creating folder.\n", outputFolder)
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
