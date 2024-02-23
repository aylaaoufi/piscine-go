package main

import "fmt"

func Join(strs []string, sep string) string {
	result := ""

	for i, char := range strs {
		if i > 0 {
			result += sep
		}
		result += char
	}
	return string(result)
}

func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(Join(toConcat, ":"))
}
