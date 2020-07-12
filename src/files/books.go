package files

import (
	"os"
	"path/filepath"
	"strings"
)

//Books list
type Books struct {
	List []bookData
}
type bookData struct {
	Name string
	File string
}

func getFilesList() []bookData {
	var files []bookData
	filepath.Walk("./files", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, bookData{Name: strings.TrimRight(info.Name(), ".epub"), File: path})
		}
		return nil
	})
	return files
}

//Init books list
func (b *Books) Init() {
	b.List = getFilesList()
}
