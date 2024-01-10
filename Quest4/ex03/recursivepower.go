package main

import "fmt"


func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 0
	}
	if power == 1 {
		return nb
	} else {
		return nb * RecursivePower(nb, power - 1)
	}
}

func main() {
	fmt.Println(RecursivePower(4, 3))
}