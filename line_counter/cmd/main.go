package main

import (
	"fmt"
	"lcounter"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		counter, err := lcounter.NewCounter(lcounter.WithInputFromArgs(os.Args[1:]))
		if err != nil {
			panic(err)
		}
		fmt.Println(counter.Count())
	} else {
		lcounter.Main()
	}
}
