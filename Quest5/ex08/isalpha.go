package main

import "fmt"

func IsAlpha(s string) bool {
	letre := []rune(s)

	for _, char := range letre {
		if !(char > 32 && char < 126) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsAlpha("Hello! How are you?"))
	fmt.Println(IsAlpha("HelloHowareyou"))
	fmt.Println(IsAlpha("What's this 4?"))
	fmt.Println(IsAlpha("Whatsthis4"))

}
