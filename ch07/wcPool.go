package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"

	"golang.org/x/sync/semaphore"
)

var data []counter

func lineByLinePool(file string, i int, sem *semaphore.Weighted) {
	defer sem.Release(1)

	lineCount := 0
	defer func() {
		data[i] = counter{kind: Line, count: lineCount, file: file}
	}()

	f, err := os.Open(file)
	if err != nil {
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		_, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			return
		}
		lineCount += 1
	}
}

func wordByWordPool(file string, i int, sem *semaphore.Weighted) {
	defer sem.Release(1)

	wordCount := 0
	defer func() {
		data[i] = counter{kind: Word, count: wordCount, file: file}
	}()

	f, err := os.Open(file)
	if err != nil {
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			return
		}

		re := regexp.MustCompile("[^\\s]+")
		words := re.FindAllString(line, -1)
		wordCount += len(words)
	}
}

func charByCharPool(file string, i int, sem *semaphore.Weighted) {
	defer sem.Release(1)

	charCount := 0
	defer func() {
		data[i] = counter{kind: Char, count: charCount, file: file}
	}()

	f, err := os.Open(file)
	if err != nil {
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			return
		}

		charCount += len(line)
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("usage: wcPool workers <file1> [<file2> ...]\n")
		return
	}

	nWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	files := os.Args[2:]
	counters := 3
	var sem = semaphore.NewWeighted(int64(nWorkers))

	bufLen := len(files) * counters
	data = make([]counter, bufLen)

	// Needed by Acquire()
	ctx := context.TODO()
	for i, file := range files {
		accSem(&ctx, sem, 1)
		go lineByLinePool(file, i*counters, sem)

		accSem(&ctx, sem, 1)
		go wordByWordPool(file, i*counters+1, sem)

		accSem(&ctx, sem, 1)
		go charByCharPool(file, i*counters+2, sem)
	}

	accSem(&ctx, sem, int64(nWorkers))

	results := make(map[string]*fileResult, len(files))
	for _, ctr := range data {
		if _, ok := results[ctr.file]; !ok {
			results[ctr.file] = &fileResult{file: ctr.file}
		}
		res := results[ctr.file]

		switch ctr.kind {
		case Line:
			res.lineCounter = ctr
		case Word:
			res.wordCounter = ctr
		case Char:
			res.charCounter = ctr
		}
	}

	printResult(files, results)
}
