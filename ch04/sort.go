package main

import (
	"fmt"
	"sort"
)

type S1 struct {
	F1 int
	F2 string
	F3 int
}

type S1slice []S1

func (a S1slice) Len() int {
	return len(a)
}

func (a S1slice) Less(i, j int) bool {
	return a[i].F3 < a[j].F3
}

func (a S1slice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	data := []S1{
		{1, "S1_1", 10},
		{2, "S1_1", 20},
		{-1, "S1_1", -20},
	}
	fmt.Println("Before:", data)
	sort.Sort(S1slice(data))
	fmt.Println("After:", data)

	// Reverse sorting works automatically
	sort.Sort(sort.Reverse(S1slice(data)))
	fmt.Println("Reverse:", data)
}
