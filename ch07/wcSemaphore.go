package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"regexp"

	"golang.org/x/sync/semaphore"
)

var buffer []counter

func lineByLineSem(file string, i int, sem *semaphore.Weighted) {
	defer sem.Release(1)

	lineCount := 0
	defer func() {
		buffer[i] = counter{kind: Line, count: lineCount, file: file}
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

func wordByWordSem(file string, i int, sem *semaphore.Weighted) {
	defer sem.Release(1)

	wordCount := 0
	defer func() {
		buffer[i] = counter{kind: Word, count: wordCount, file: file}
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

func charByCharSem(file string, i int, sem *semaphore.Weighted) {
	defer sem.Release(1)

	charCount := 0
	defer func() {
		buffer[i] = counter{kind: Char, count: charCount, file: file}
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
	if len(os.Args) == 1 {
		fmt.Printf("usage: wcSemaphone <file1> [<file2> ...]\n")
		return
	}

	files := os.Args[1:]
	counters := 3
	var sem = semaphore.NewWeighted(int64(counters))

	bufLen := len(files) * counters
	buffer = make([]counter, bufLen)

	// Needed by Acquire()
	ctx := context.TODO()
	for i, file := range files {
		accSem(&ctx, sem, 1)
		go lineByLineSem(file, i*counters, sem)

		accSem(&ctx, sem, 1)
		go wordByWordSem(file, i*counters+1, sem)

		accSem(&ctx, sem, 1)
		go charByCharSem(file, i*counters+2, sem)
	}

	accSem(&ctx, sem, int64(counters))

	results := make(map[string]*fileResult, len(files))
	for _, ctr := range buffer {
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
