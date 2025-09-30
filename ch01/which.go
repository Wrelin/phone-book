package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
		return
	}

	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	for _, directory := range pathSplit {
		for _, file := range arguments[1:] {
			fullPath := filepath.Join(directory, file)
			checkAndPrint(fullPath)
		}
	}
}

func checkAndPrint(fullPath string) {
	// Does it exist?
	fileInfo, err := os.Stat(fullPath)
	if err != nil {
		return
	}

	mode := fileInfo.Mode()
	// Is it a regular file?
	if !mode.IsRegular() {
		return
	}

	// Is it executable?
	if mode&0111 != 0 {
		fmt.Println(fullPath)
	}
}
