package main

import "fmt"

// Add two number
func Add(a int, b int) int {
	return a + b
}

// Operation1 on number
func Operation1(a int, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b
}

// Operation on number
func Operation(a int, b int) []int {
	return []int{a + b, a - b, a * b, a / b}
}

// MulOperation paramenter
func MulOperation(Numbers ...int) int {
	total := 0
	for i := 0; i < len(Numbers); i++ {
		total += Numbers[i]
	}
	return total
}

// myRecursion: Recursion function for factorial
func myRecursion(number int) int {
	if number == 1 {
		return 1
	}
	return number * myRecursion(number-1)
}

// Calculate number
func Calculate() {
	fmt.Println(Add(5, 3))
	fmt.Println(Operation(5, 3))
	fmt.Println(Operation1(5, 3))
	ad, _, mul, _ := Operation1(33, 4) // _ is placeholder in go
	fmt.Println(ad, mul)
	fmt.Println(MulOperation(11, 22, 33, 44, 55))
	Numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(MulOperation(Numbers...))

	//Factorial
	fmt.Println("Factorial 0f 5 : ", myRecursion(5))
}
