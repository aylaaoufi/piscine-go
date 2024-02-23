package main

import (
	"fmt"
)

func StrRev(s string) string {
	rev := []rune(s)
	i := len(rev) - 1
	j := 0

	for j < i {
		rev[i], rev[j] = rev[j], rev[i]
		i--
		j++
	}
	return string(rev)
}

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}
