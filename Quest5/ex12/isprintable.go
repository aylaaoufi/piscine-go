package main

import "fmt"

func IsPrintable(s string) bool {
	letre := []rune(s)

	for _, char := range letre {
		if !(char >= 32 && char <= 126) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPrintable("Hello"))
	fmt.Println(IsPrintable("Hello\n"))
}
