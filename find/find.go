package find

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func Files(fsys fs.FS) (files []string) {
	// fsys := os.DirFS(path)
	fs.WalkDir(fsys, ".", func(path string, dir fs.DirEntry, err error) error {
		if err != nil {
			return fs.SkipDir
		}
		if filepath.Ext(path) == ".go" {
			files = append(files, path)
		}
		return nil
	})
	return files
}

func Main() {
	fsys := os.DirFS(".")
	// f, err := fsys.Open("file.go")
	var count int
	fs.WalkDir(fsys, ".", func(path string, dir fs.DirEntry, err error) error {
		if err != nil {
			return fs.SkipDir
		}
		if filepath.Ext(path) == ".go" {
			count++
		}
		return nil

	})
	fmt.Println(count)
}
