package main

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func main() {
	in := os.Args[1]
	out := os.Args[2]
	outFile, err := os.Create(out)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	writer := zip.NewWriter(outFile)
	defer writer.Close()
	filepath.Walk(in, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		header, err := zip.FileInfoHeader(info)
		fp, err := filepath.Rel(in, path)
		header.Name = filepath.ToSlash(in + "/" + fp)
		if err != nil {
			panic(err)
		}
		if !info.IsDir() {
			fileWriter, err := writer.CreateHeader(header)
			header.Method = zip.Deflate
			header.Method = zip.Deflate
			file, err := os.Open(path)
			if err != nil {
				panic(err)
			}
			defer file.Close()
			io.Copy(fileWriter, file)
		}
		return nil
	})
}
