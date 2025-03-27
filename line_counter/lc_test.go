package lcounter_test

import (
	"bytes"
	"lcounter"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Parallel()
	// in := new(bytes.Buffer)
	// fmt.Fprint(in, "line1\nline2\nline3")
	in := bytes.NewBufferString("line1\nline2\nline3")
	c, err := lcounter.NewCounter(lcounter.WithInput(in))
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Count()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}
