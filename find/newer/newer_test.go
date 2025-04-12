package newer_test

import (
	"newer"
	"testing"
	"testing/fstest"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestNew_FindsRecentFiles(t *testing.T) {
	t.Parallel()
	now := time.Now()
	files := fstest.MapFS{
		"file0":                     {ModTime: now.Add(-30 * 24 * time.Hour)},
		"file1":                     {ModTime: now},
		"folder/file2":              {ModTime: now.Add(-30 * 24 * time.Hour)},
		"folder/subdirectory/file3": {ModTime: now},
	}
	want := []string{
		"file1",
		"folder/subdirectory/file3",
	}
	got := newer.Files(files, 30)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
