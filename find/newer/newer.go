package newer

import (
	"io/fs"
	"time"
)

func Files(filesys fs.FS, days int) (files []string) {
	threshold := time.Now().Add(-24 * time.Hour * time.Duration(days))
	fs.WalkDir(filesys, ".", func(path string, entry fs.DirEntry, err error) error {
		info, err := entry.Info()
		if err != nil || info.IsDir() {
			return nil
		}
		if info.ModTime().After(threshold) {
			files = append(files, path)
		}
		return nil
	})
	return files
}
