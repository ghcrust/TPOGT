package hello

import (
	"fmt"
	"io"
)

func SayHelloNameToWriter(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}
