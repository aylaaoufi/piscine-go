package main

import "fmt"


func IsUpper(s string) bool {
	letre := []rune(s)

	for _, char := range letre {
		if !(char >= 'A' && char <= 'Z') {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsUpper("HELLO"))
	fmt.Println(IsUpper("HELLO!"))
}