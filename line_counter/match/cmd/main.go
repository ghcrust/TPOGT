package main

import (
	"fmt"
	"match"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: match <string_to_match> <files> <to> <search>")
		os.Exit(1)
	}
	for _, arg := range os.Args[2:] {
		file, err := os.Open(arg)
		if err != nil {
			panic(err)
		}
		m := match.NewMatcher().WithReader(file)
		m.Match(os.Args[1])
	}
}
