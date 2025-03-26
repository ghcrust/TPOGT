package hello_test

import (
	"bytes"
	"hello"
	"testing"
)

func TestPrintHelloToWriter(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	hello.SayHelloNameToWriter(buf, "x")
	want := "Hello, x"
	got := buf.String()
	if want != got {
		t.Errorf("want %q got %q", want, got)
	}
}
