package main

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/sync/semaphore"
)

func printResult(files []string, results map[string]*fileResult) {
	totalLines := 0
	totalWords := 0
	totalChars := 0
	for _, file := range files {
		res := results[file]
		totalLines += res.lineCounter.count
		totalWords += res.wordCounter.count
		totalChars += res.charCounter.count
		fmt.Println(res.lineCounter.count, res.wordCounter.count, res.charCounter.count, file)
	}

	if len(files) > 1 {
		fmt.Println(totalLines, totalWords, totalChars, "total")
	}
}

func accSem(ctx *context.Context, sem *semaphore.Weighted, weight int64) {
	err := sem.Acquire(*ctx, weight)
	if err != nil {
		fmt.Println("Cannot acquire semaphore:", err)
		os.Exit(1)
	}
}
