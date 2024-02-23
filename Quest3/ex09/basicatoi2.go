package main

import "fmt"

func BasicAtoi2(s string) int {
	result := 0
	for _, char := range s {
		result = result*10 + int(char-'0')
		if char == ' ' || char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			return 0
		}
	}
	return result
}

func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}
