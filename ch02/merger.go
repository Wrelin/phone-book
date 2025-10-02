package main

import "fmt"

const (
	lenA   = 3
	lenB   = 4
	lenRes = lenA + lenB
)

func mergeArraysToSlice[T any](arr1 [lenA]T, arr2 [lenB]T) []T {
	res := make([]T, len(arr1), len(arr1)+len(arr2))
	copy(res, arr1[:])
	return append(res, arr2[:]...)
}

func mergeArraysToArray[T any](arr1 [lenA]T, arr2 [lenB]T) [lenRes]T {
	res := [lenRes]T{}
	for i, n := range arr1 {
		res[i] = n
	}
	for i, n := range arr2 {
		res[i+len(arr1)] = n
	}
	return res
}

func mergeSlicesToArray[T any](arr1 []T, arr2 []T) [lenRes]T {
	res := [lenRes]T{}
	for i, n := range arr1 {
		if i >= lenRes {
			return res
		}
		res[i] = n
	}

	for i, n := range arr2 {
		if i+len(arr1) >= lenRes {
			return res
		}
		res[i+len(arr1)] = n
	}

	return res
}

func main() {
	arr1 := [...]int{1, 2, 3}
	arr2 := [...]int{4, 5, 6, 7}

	fmt.Printf("%+v\n", mergeArraysToSlice(arr1, arr2))
	fmt.Printf("%+v\n", mergeArraysToArray(arr1, arr2))
	fmt.Printf("%+v\n", mergeSlicesToArray(arr1[:], arr2[:]))
}
