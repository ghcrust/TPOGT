package writer_test

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"writer"

	"github.com/google/go-cmp/cmp"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	testscript.Main(m, map[string]func(){
		"writefile": func() {
			writer.Main()
		},
		"mkfile": func() {
			if len(os.Args) < 3 {
				fmt.Fprintln(os.Stderr, "Usage: mkfile <size> <file_path>")
				os.Exit(1)
			}
			size, err := strconv.Atoi(os.Args[1])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Invalid size: %s", os.Args[1])
				os.Exit(1)
			}

			file := os.Args[2]
			err = os.WriteFile(file, make([]byte, size), 0o666)
			if err != nil {
				fmt.Fprintf(os.Stderr, "WriteFile: %s", err)
				os.Exit(1)
			}
		},
	})
}

func Test(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}

func TestWriteToFile_WritesDataToFile(t *testing.T) {
	t.Parallel()
	/*
		path := "testdata/write_test.txt"
		if _, err := os.Stat(path); err == nil {
			t.Fatalf("test artifact not cleaned up: %q", path)
		}
		defer os.Remove(path)
	*/
	path := os.TempDir() + "write_test.txt"
	data := []byte("test\ndata")
	err := writer.WriteToFile(path, data)
	if err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(data, got) {
		t.Fatal(cmp.Diff(data, got))
	}
	stat, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	perm := stat.Mode().Perm()
	if perm != 0o666 {
		t.Errorf("want file mode 0o666 got O%o", perm)
	}
}

func TestWriteToFile_ReturnsErrorUponUnwritablePath(t *testing.T) {
	t.Parallel()
	path := "unwriteable/path/file.txt"
	err := writer.WriteToFile(path, []byte("..."))
	if err == nil {
		t.Fatal("want error for invalid file path, got nil")
	}
}

func TestWriteToFile_OverwritesExistingFiles(t *testing.T) {
	t.Parallel()
	path := os.TempDir() + "file_test_overwrite.txt"
	err := os.WriteFile(path, []byte("test"), 0o666)
	if err != nil {
		t.Fatal(err)
	}
	want := []byte("test_overwrite")
	err = writer.WriteToFile(path, want)
	if err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestWriteToFile_ChangesPermsOnExistingFile(t *testing.T) {
	t.Parallel()
	path := os.TempDir() + "file_test_permissions.txt"
	err := os.WriteFile(path, []byte("test"), 0o644)
	if err != nil {
		t.Fatal(err)
	}
	err = writer.WriteToFile(path, []byte("overwrite"))
	if err != nil {
		t.Fatal(err)
	}
	stat, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	perm := stat.Mode().Perm()
	if perm != 0o666 {
		t.Errorf("want file mode 0o666 got 0o%o", perm)
	}
}

func TestWriteBytes_WritesZerosToFile(t *testing.T) {
	t.Parallel()

}
