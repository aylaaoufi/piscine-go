package main

import (
	"os"
	"github.com/01-edu/z01"
)


func main() {
	input := os.Args[0]

	for _, char := range input {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}