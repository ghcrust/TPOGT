package match_test

import (
	"bytes"
	"match"
	"testing"
)

func TestMatch(t *testing.T) {
	in := bytes.NewBufferString("test\nline2\nline 3\nmatch this line\nline5")
	out := new(bytes.Buffer)
	m := match.NewMatcher().WithReader(in).WithWriter(out)
	want := "match this line\n"
	m.Match("match")
	got := out.String()
	if want != got {
		t.Errorf("want %q got %q", want, got)
	}
}
