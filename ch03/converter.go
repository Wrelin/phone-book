package main

import "fmt"

func toMap[T comparable](arr [3]T) map[T]T {
	res := map[T]T{}
	for _, val := range arr {
		res[val] = val
	}
	return res
}

func toSlices[T comparable, V any](data map[T]V) ([]T, []V) {
	keys := make([]T, 0, len(data))
	vals := make([]V, 0, len(data))

	for key, val := range data {
		keys = append(keys, key)
		vals = append(vals, val)
	}

	return keys, vals
}

func main() {
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]string{"a", "b", "c"}

	fmt.Printf("%+v\n", toMap(arr1))
	fmt.Printf("%+v\n", toMap(arr2))

	data := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	keys, vals := toSlices(data)
	fmt.Printf("%+v, %+v\n", keys, vals)
}
