package main

import "fmt"

func Sqrt(nb int) int {
	for i := 2; i*i <= nb; i++ {
		if i*i == nb {
			return i
		}
	}
	return 0
}

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(3))
}
