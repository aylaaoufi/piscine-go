package main

import "fmt"

func ToUpper(s string) string {
	letre := []rune(s)
	for i, char := range letre {
		if char >= 'a' && char <= 'z' {
			letre[i] -= 32
		}
	}
	return string(letre)
}

func main() {
	fmt.Println(ToUpper("Hello! How are you?"))
}
