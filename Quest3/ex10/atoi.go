package main

import "fmt"

func Atoi(s string) int {
	result := 0
	sign := 1

	for i, char := range s {
		if i == 0 && (char == '-' || char == '+') {
			if char == '-' {
				sign = -sign
			}
			continue
		}
		if char < '0' || char > '9' {
			return 0
		}
		result = result*10 + int(char-'0')
	}
	return result * sign
}

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}
