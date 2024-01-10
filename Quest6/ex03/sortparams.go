package main

import (
	"os"

	"fmt"
)


func main() {
	input := os.Args[1:] 

	for i := 0; len(input) > i; i++ {
		for j := i + 1; len(input) > j; j++ {
			if input[i] > input[j] {
				input[i], input[j] = input[j], input[i]
			}
		}
	}
	for _, arg := range input {
		fmt.Println(arg)
	}
}