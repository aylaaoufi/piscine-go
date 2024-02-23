package main

import "fmt"

func UltimateDivMod(a *int, b *int) {
	x := *a
	*a = x / *b
	*b = x % *b
}

func main() {
	a := 13
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
