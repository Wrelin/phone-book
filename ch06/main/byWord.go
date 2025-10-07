package main

import (
	"bufio"
	"fmt"
	"os"
)

func wordByWord(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return scanner.Err()
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("usage: byWord <file1> [<file2> ...]\n")
		return
	}

	for _, file := range args[1:] {
		err := wordByWord(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}
