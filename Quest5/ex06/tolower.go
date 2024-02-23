package main

import "fmt"

func ToLower(s string) string {
	letre := []rune(s)

	for i, char := range letre {
		if char >= 'A' && char <= 'Z' {
			letre[i] += 32
		}
	}
	return string(letre)
}

func main() {
	fmt.Println(ToLower("Hello! How are you?"))
}
