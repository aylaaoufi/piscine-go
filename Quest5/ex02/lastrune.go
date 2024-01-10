package main

import "github.com/01-edu/z01"


func LastRune(s string) rune {
	letre := []rune(s)
	length := len(letre)
	return letre[length-1]
}

func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("Ola!"))
	z01.PrintRune('\n')
}