package main

import (
	"fmt"
	"os"
)

type Argument struct {
	index int
	value string
}

func main() {
	arguments := make([]Argument, len(os.Args))
	for i, arg := range os.Args {
		arguments[i] = Argument{
			index: i,
			value: arg,
		}
	}

	fmt.Printf("%+v\n", arguments)
}
