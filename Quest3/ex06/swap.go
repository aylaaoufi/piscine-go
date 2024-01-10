package main

import "fmt"


func Swap(a *int, b *int) {
	x := *a
	*a = *b
	*b = x
}

func main() {
	a := 0
	b := 1
	Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}