package find_test

import (
	"archive/zip"
	"find"
	"os"
	"testing"
	"testing/fstest"

	"github.com/google/go-cmp/cmp"
)

func TestFindFiles_ListsFilesInTree(t *testing.T) {
	t.Parallel()
	want := []string{
		"file.go",
		"subfolder/file1.go",
		"subfolder2/file2.go",
		"subfolder2/file3.go",
	}
	got := find.Files(os.DirFS("testdata/tree"))
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestFindFiles_ListsFilesInMapFS(t *testing.T) {
	t.Parallel()
	fsys := fstest.MapFS{
		"file.go":             {},
		"subfolder/file1.go":  {},
		"subfolder2/file2.go": {},
		"subfolder2/file3.go": {},
	}
	want := []string{
		"file.go",
		"subfolder/file1.go",
		"subfolder2/file2.go",
		"subfolder2/file3.go",
	}
	got := find.Files(fsys)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestFind_ListsFilesInZipArchive(t *testing.T) {
	t.Parallel()
	fsys, err := zip.OpenReader("testdata/tree.zip")
	if err != nil {
		t.Fatal(err)
	}
	want := []string{
		"tree/file.go",
		"tree/subfolder/file1.go",
		"tree/subfolder2/file2.go",
		"tree/subfolder2/file3.go",
	}
	got := find.Files(fsys)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func BenchmarkFilesOnDisk(b *testing.B) {
	fsys := os.DirFS("testdata/tree")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = find.Files(fsys)
	}
}

func BenchmarkFilesInMemory(b *testing.B) {
	fsys := fstest.MapFS{
		"file.go":             {},
		"subfolder/file1.go":  {},
		"subfolder2/file2.go": {},
		"subfolder2/file3.go": {},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		find.Files(fsys)
	}
}
