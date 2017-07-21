package proc

import (
	"os"
	"fmt"
	"io/ioutil"
	"strings"
	"path/filepath"
)

type FileInfoWithPath struct {
	Path         string
	Folder       string
	RelativePath string
	Name         string
	IsDir        bool
}

func (f FileInfoWithPath) String() string {
	return fmt.Sprintf("path: %s, folder: %s, isDir: %t", f.Path, f.Folder, f.IsDir)
}

func lsOSFileInfo(dir string) []os.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	return files
}

func Ls(dir string, filter func(string) bool, verbose bool) []FileInfoWithPath {
	info := lsInfo(dir, "./", filter, verbose)
	return info
}

func getFolderNameFromDir(dir string) string {
	folder := dir
	// todo windows .\ ?
	if dir == "." || dir == "./" {
		var err error
		folder, err = filepath.Abs(".")
		if err != nil {
			panic(err)
		}
	}

	f := strings.Split(folder, "/")
	// Windows Path g:\dir\dir
	if len(f) == 1 {
		f = strings.Split(folder, "\\")
	}
	folder = f[len(f)-1]
	return folder
}

func lsInfo(dir string, rp string, filter func(string) bool, verbose bool) []FileInfoWithPath {
	folder := getFolderNameFromDir(dir)
	return lsInfoInner(dir, folder, rp, filter, verbose)
}

func lsInfoInner(dir, folder string, relativePath string, filter func(string) bool, verbose bool) []FileInfoWithPath {
	var filesInfo []FileInfoWithPath
	files := lsOSFileInfo(dir)

	// ignore "./"
	if dir == "." || dir == "./" {
		dir = ""
	} else {
		if dir[len(dir)-1:] != "/" {
			dir = dir + "/"
		}
	}

	if relativePath == "." || relativePath == "./" {
		relativePath = ""
	}

	for _, f := range files {
		if f.IsDir() {
			filesInfo = append(filesInfo, lsInfoInner(dir+f.Name(), f.Name(), relativePath+"/"+f.Name(), filter, verbose)...)
		} else {
			if filter(f.Name()) {
				filesInfo = append(filesInfo, FileInfoWithPath{dir /*+ f.Name()*/ , folder, relativePath, f.Name(), f.IsDir()})
			}
		}
	}
	return filesInfo
}
