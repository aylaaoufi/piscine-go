package main

import "fmt"


func IterativePower(nb int, power int) int {
	result := 1

	if power < 0 {
		return 0
	}
	for i := 1; i <= power; i++ {
		result *= nb
	}
	return result
}

func main() {
	fmt.Println(IterativePower(4, 3))
}