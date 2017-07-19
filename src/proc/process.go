package proc

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

func getFileName(s string) string {
	var extension = filepath.Ext(s)
	return s[0:len(s)-len(extension)]
}

func generateContent(info FileInfoWithPath) string {
	content, _ := ioutil.ReadFile(info.path)
	output := "package " + info.folder +
		"\n\nconst " + strings.Title(getFileName(info.info.Name())) + " = \"" + strings.TrimSpace(string(content)) + "\""
	return output
}

func generateGoFile(info []FileInfoWithPath) []string{
	var result []string
	for _, i := range info {
		content := generateContent(i)
		// todo output dir
		err := ioutil.WriteFile(i.path + ".go", []byte(content), 0744)
		if err != nil {
			panic(err)
		}
	}
	return result
}

func FossilDir(dir string){
	// todo flags
	info := lsInfo(dir)
	generateGoFile(info)
}