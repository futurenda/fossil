package proc

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"fmt"
	"os"
	"strconv"
	"gopkg.in/cheggaaa/pb.v1"
	"time"
	"sync"
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

func generateContent(info FileInfoWithPath, paras Paras) string {
	content, err := ioutil.ReadFile(info.Path + info.Name)
	if err != nil {
		panic(err)
	}

	sqlVarName := snakeToCamelCase(regularizeToSnakeCase(getFileName(info.Name)))
	sqlContent := processContents([]byte(strings.TrimSpace(string(content))), false)
	packageName := info.Folder
	if paras.Package != "" {
		packageName = paras.Package
	}
	output := fmt.Sprintf("package %s\n\nconst %s = %s\n",
		packageName, sqlVarName, sqlContent)
	return output
}

func generateGoFile(i FileInfoWithPath, paras Paras) {
	content := generateContent(i, paras)
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
	bar := pb.New(len(info)).Prefix("Progress ")
	barPool, err := pb.StartPool(bar)
	if err != nil {
		panic(err)
	}
	chans := make(chan bool, paras.Limit)
	wg := new(sync.WaitGroup)

	for _, i := range info {
		chans <- true
		wg.Add(1)
		go func(info FileInfoWithPath, bar *pb.ProgressBar) {
			generateGoFile(info, paras)
			bar.Increment()
			time.Sleep(5 * time.Second)
			<-chans
			wg.Done()
		}(i, bar)
	}
	wg.Wait()

	barPool.Stop()
}

type Paras struct {
	InputPath  string
	OutputPath string
	Verbose    bool
	Limit      int
	Package    string
}

type FossilInfo struct {
	Count int
}

func FossilDir(paras Paras) FossilInfo {
	if paras.Verbose{
		fmt.Printf("Input from %s, output to %s, limit %d.", paras.InputPath, paras.OutputPath, paras.Limit)
	}

	if paras.OutputPath == "" {
		paras.OutputPath = paras.InputPath
	}

	sqlFileFilter := func(s string) bool {
		return filepath.Ext(s) == ".sql"
	}
	info := Ls(paras.InputPath, sqlFileFilter, paras.Verbose)
	if paras.Verbose {
		fmt.Printf("Found %d files\n", len(info))
	}
	generateAllFile(info, paras)
	return FossilInfo{len(info)}
}
