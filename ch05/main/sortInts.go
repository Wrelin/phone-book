package main

import (
	"fmt"
	"slices"
)

func sortUnnamed(i, j, k int) (int, int, int) {
	s := []int{i, j, k}
	slices.Sort(s)
	return s[0], s[1], s[2]
}

func sortNamed(i, j, k int) (min, mid, max int) {
	s := []int{i, j, k}
	slices.Sort(s)
	min, mid, max = s[0], s[1], s[2]
	return
}

func main() {
	fmt.Println(sortUnnamed(3, 1, 2))
	fmt.Println(sortNamed(3, 1, 2))
}
