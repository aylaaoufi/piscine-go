package main

import "fmt"

func Capitalize(s string) string {
	letre := []rune(s)
	nextchar := true

	for i, char := range letre {
		if Capcase(char) {
			if nextchar {
				letre[i] = Toupperr(char)
				nextchar = false
			}
		} else {
			nextchar = true
		}
	}
	return string(letre)
}

func Capcase(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func Toupperr(char rune) rune {
	if char >= 'a' && char <= 'z' {
		char -= 32
	}
	return char
}

func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}
