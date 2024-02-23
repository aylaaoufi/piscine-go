package main

import "fmt"

func SortIntegerTable(table []int) {
	for i := 0; len(table) > i; i++ {
		for j := i + 1; len(table) > j; j++ {
			if table[i] > table[j] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
}

func main() {
	s := []int{5, 4, 3, 2, 1, 0}
	SortIntegerTable(s)
	fmt.Println(s)
}
