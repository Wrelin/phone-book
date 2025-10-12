package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sync"
)

func lineByLineCh(file string, ch chan<- counter, wg *sync.WaitGroup) {
	defer wg.Done()

	lineCount := 0
	defer func() {
		ch <- counter{kind: Line, count: lineCount, file: file}
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

func wordByWordCh(file string, ch chan<- counter, wg *sync.WaitGroup) {
	defer wg.Done()

	wordCount := 0
	defer func() {
		ch <- counter{kind: Word, count: wordCount, file: file}
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

func charByCharCh(file string, ch chan<- counter, wg *sync.WaitGroup) {
	defer wg.Done()

	charCount := 0
	defer func() {
		ch <- counter{kind: Char, count: charCount, file: file}
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
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("usage: wcChan <file1> [<file2> ...]\n")
		return
	}

	files := args[1:]
	counters := 3

	bufLen := len(files) * counters
	ch := make(chan counter, bufLen)
	var wg sync.WaitGroup
	for _, file := range files {
		wg.Add(counters)
		go lineByLineCh(file, ch, &wg)
		go wordByWordCh(file, ch, &wg)
		go charByCharCh(file, ch, &wg)
	}

	wg.Wait()

	results := make(map[string]*fileResult, len(files))
	bufReaded := 0
	for ctr := range ch {
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

		bufReaded++
		if bufReaded == bufLen {
			close(ch)
		}
	}

	printResult(files, results)
}
