package main

type counterKind int

const (
	Char counterKind = iota
	Word
	Line
)

type counter struct {
	kind  counterKind
	count int
	file  string
}

type fileResult struct {
	file        string
	lineCounter counter
	wordCounter counter
	charCounter counter
}
