package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchNameSur(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
	return re.Match(t)
}

func main() {
	arguments := os.Args[1:]
	if len(arguments) == 0 {
		fmt.Println("Usage: <utility1> <utility2>...")
		return
	}

	for _, arg := range arguments {
		ret := matchNameSur(arg)
		fmt.Println(ret)
	}
}
