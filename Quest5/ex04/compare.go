package main

import "fmt"

func Compare(a, b string) int {
	a1 := []rune(a)
	b1 := []rune(b)

	for i := 0; len(a1) > i && len(b1) > i; i++ {
		if a1[i] < b1[i] {
			return -1
		} else if a1[i] > b1[i] {
			return 1
		}
		if len(a1) < len(b1) {
			return -1
		} else if len(a1) > len(b1) {
			return 1
		}
	}
	return 0
}

func main() {
	fmt.Println(Compare("Hello!", "Hello!"))
	fmt.Println(Compare("Salut!", "lut!"))
	fmt.Println(Compare("Ola!", "Ol"))
}
