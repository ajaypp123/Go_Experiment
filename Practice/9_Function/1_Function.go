/*
Function:
	- function can be global or package level.
		Global function added as Capital letter
		Package Function can be started with small letter
	- Passing pointer is good than passing value
	- Return local variable from func
		- In any other language returning value from function can cause segmentation fault.
		But go will move that variable to heap and return its value.
		Example: func fun() int {
					c:= 44 ; return c
				}

	- Declare return variable in function definition
		func fun() (result int) {
			result = 22
			return // it will return result variable
		}

	- Function can return error object that can be used to handle code
		func myError() (error) {
			return fmt.Errorf("Error code")
		}
*/

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

// funPointer return pointer to function
func funPointer(a int, b int) *int {
	c := a + b
	return &c
}

func funReturn() (result int) {
	result = 22
	return // it will return result variable
}

func handleError(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Can not divided by zero")
	}
	return a / b, nil
}

// Anonymous function

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
	var c *int = funPointer(22, 33)
	fmt.Println(*c)
	fmt.Println(funReturn())

	fmt.Println(handleError(7, 0))
	fmt.Println(handleError(7, 3))
}

/*
func main() {
	Calculate()
}
*/
