package main

import "fmt"


func BasicJoin(elems []string) string {
	result := ""

	for _, char := range elems {
		result += char
	}
	return result
}

func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(BasicJoin(elems))
}