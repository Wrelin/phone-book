package main

import (
	"fmt"
	"reflect"
)

type s1 struct {
	i int
}

type s2 struct {
	s string
}

func getKind(s any) string {
	switch t := s.(type) {
	case s1:
		return "s1"
	case s2:
		return "s2"
	default:
		return reflect.TypeOf(t).String()
	}
}

func main() {
	fmt.Println(getKind(s1{}))
	fmt.Println(getKind(s2{}))
	fmt.Println(getKind(123))
	fmt.Println(getKind(123.0))
	fmt.Println(getKind("123"))
}
