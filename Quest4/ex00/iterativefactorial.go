package main

import "fmt"

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	result := 1
	for i := 1; i <= nb; i++ {
		result *= i
	}
	return result
}

func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}
