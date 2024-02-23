package main

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	PrintDigits(n)
}

func PrintDigits(n int) {
	if n < 10 {
		digit := n % 10
		z01.PrintRune(rune(digit + '0'))
	} else {
		PrintDigits(n / 10)
		digit := n % 10
		z01.PrintRune(rune(digit + '0'))
	}
}

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune('\n')
}
