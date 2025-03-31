package hello

import (
	"fmt"
	"io"
	"os"
)

type Printer struct {
	DefaultWriter io.Writer
}

func NewPrinter() *Printer {
	return &Printer{os.Stdout}
}

func (p *Printer) PrintHelloName(name string) {
	fmt.Fprintf(p.DefaultWriter, "Hello, %s", name)
}

func Main() {
	NewPrinter().PrintHelloName("world")
}
