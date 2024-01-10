package main

import "fmt"


func IsPrime(nb int) bool {
	for i := 2; i * i <= nb; i++ {
		if nb % i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPrime(5))
	fmt.Println(IsPrime(4))
}