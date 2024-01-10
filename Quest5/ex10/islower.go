package main

import "fmt"


func IsLower(s string) bool {
	letre := []rune(s)

	for _, char := range letre {
		if !(char >= 'a' && char <= 'z') {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsLower("hello"))
	fmt.Println(IsLower("hello!"))

}