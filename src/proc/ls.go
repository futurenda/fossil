package proc

import (
	"os"
	"fmt"
	"io/ioutil"
	"strings"
	"path/filepath"
)

type FileInfoWithPath struct {
	path   string
	folder string
	info   os.FileInfo
}

func (f FileInfoWithPath) String() string {
	return fmt.Sprintf("path: %s, folder: %s, isDir: %t", f.path, f.folder, f.info.IsDir())
}

func ls(dir string) []os.FileInfo {
	files, _ := ioutil.ReadDir(dir)
	return files
}

func Ls(dir string) []FileInfoWithPath {
	info := lsInfo(dir)
	return info
}

func lsInfo(dir string) []FileInfoWithPath {
	folder := getFolderNameFromDir(dir)
	return lsInfoInner(dir, folder)
}

func getFolderNameFromDir(dir string) string{
	folder := dir
	// todo windows .\ ?
	if dir == "." || dir == "./" {
		folder, _ = filepath.Abs(".")
	}

	f := strings.Split(dir, "/")
	// Windows Path g:\dir\dir
	if len(f) == 1 {
		f = strings.Split(dir, "\\")
	}
	folder = f[len(f)-1]
	return folder
}

func lsInfoInner(dir, folder string) []FileInfoWithPath {
	var filesInfo []FileInfoWithPath
	files := ls(dir)

	// ignore "./"
	if dir == "." || dir == "./" {
		dir = ""
	} else {
		dir = dir + "/"
	}

	for _, f := range files {
		if f.IsDir() {
			filesInfo = append(filesInfo, lsInfoInner(dir+f.Name(), f.Name())...)
		} else {
			filesInfo = append(filesInfo, FileInfoWithPath{dir + f.Name(), folder, f})
		}
	}
	return filesInfo
}
