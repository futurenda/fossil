package proc

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

func get_file_name(s string) string {
	var extension = filepath.Ext(s)
	return s[0:len(s)-len(extension)]
}

func generate(info FileInfoWithPath) string {
	content, _ := ioutil.ReadFile(info.path)
	output := "package " + info.folder +
		"\n\nconst " + strings.Title(get_file_name(info.info.Name())) + " = \"" + strings.TrimSpace(string(content)) + "\""
	return output
}

func process(info []FileInfoWithPath){
	for _, i := range info {
		content := generate(i)
		// todo output path
		err := ioutil.WriteFile(i.path + ".go", []byte(content), 0744)
		if err != nil {
			panic(err)
		}
	}
}

func FossilDir(dir string){
	info := ls_info(dir)
	process(info)
}