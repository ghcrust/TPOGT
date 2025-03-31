package hello_test

import (
	"bytes"
	"hello"
	"testing"
)

func TestPrintHelloToWriter(t *testing.T) {
	t.Parallel()
	printer := hello.NewPrinter()
	buf := new(bytes.Buffer)
	printer.DefaultWriter = buf
	printer.PrintHelloName("x")
	want := "Hello, x"
	got := buf.String()
	if want != got {
		t.Errorf("want %q got %q", want, got)
	}
}

