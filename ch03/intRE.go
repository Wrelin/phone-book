package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchInt(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[-+]?\d+$`)
	return re.Match(t)
}

func main() {
	arguments := os.Args[1:]
	if len(arguments) == 0 {
		fmt.Println("Usage: <utility1> <utility2>...")
		return
	}

	for _, arg := range arguments {
		ret := matchInt(arg)
		fmt.Println(ret)
	}
}
