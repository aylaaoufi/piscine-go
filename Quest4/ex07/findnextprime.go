package main

import "fmt"


func FindNextPrime(nb int) int {
	if nb == 0 {
		return 2
	}

	for i := 2; i * i <= nb; i++ {
		if nb % i == 0{
			return nb + 1
		} 
	}
	return nb
}

func main() {
	fmt.Println(FindNextPrime(5))
	fmt.Println(FindNextPrime(4))
}