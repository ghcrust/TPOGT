package main

import (
	"fmt"
	"lcounter"
	"log"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			file, err := os.Open(arg)
			if err != nil {
				log.Fatal(err)
			}
			counter, err := lcounter.NewCounter(lcounter.WithInput(file))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("number of lines in %v is: %v", arg, counter.Count())
		}
	} else {
		fmt.Printf("Usage: lc <files> <to> <count>\n")
	}
}
