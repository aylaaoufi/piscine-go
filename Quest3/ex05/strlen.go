package main

import "fmt"

func StrLen(s string) int {
	for range s {

	}
	return len(s)
}

func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}
