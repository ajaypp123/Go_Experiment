package main

import "fmt"

func myRange() {
	myArr := []int{1, 2, 3, 4}
	for i, c := range myArr {
		fmt.Println(i, c)
	}
	fmt.Println()

	myDict := map[string]int{"a": 1, "b": 2}
	for i, c := range myDict {
		fmt.Println(i, c)
	}
}
