package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	input := os.Args[1:]
	for i := len(input) - 1; i >= 0; i-- {
		for _, char := range input[i] {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
	z01.PrintRune(' ')
}
