package main

import "fmt"


func Index(s string, toFind string) int {
	letre := []rune(s)
	letre2 := []rune(toFind)

	for i := 0; len(letre) > i; i++ {
		for j := 0; len(letre2) > j; j++ {
			if letre[i] == letre2[j] {
				return i
			} else if letre[i] != letre2[j] {
				break
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}