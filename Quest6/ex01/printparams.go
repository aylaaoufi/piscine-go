package main

import (
	"os"
	"github.com/01-edu/z01"
)


func main() {
	input := os.Args[1:]

	for _, arg := range input {
		for _, char := range arg {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			z01.PrintRune(char)
		} else if char == ' ' {
			z01.PrintRune('\n')
		}
	  }
	  z01.PrintRune('\n')
	}
}