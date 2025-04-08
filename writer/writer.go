package writer

import (
	"flag"
	"fmt"
	"os"
)

func WriteToFile(path string, data []byte) error {
	err := os.WriteFile(path, data, 0o666)
	if err != nil {
		return err
	}
	return os.Chmod(path, 0o666)
}

func Main() int {
	size := flag.Int("size", 0, "size in bytes to write")
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "No file path to write specified")
		return 1
	}
	path := args[0]
	err := WriteToFile(path, make([]byte, *size))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to file: %s\n", err)
		return 1
	}
	return 0
}
