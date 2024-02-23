package main

import "fmt"

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	} else if index == 1 {
		return 1
	}
	if index == 0 {
		return 0
	}
	return Fibonacci(index-1) + Fibonacci(index-2)
}

func main() {
	arg1 := 4
	fmt.Println(Fibonacci(arg1))
}
