/*
Pointer
Variable that store address of other variable
*/

package main

import "fmt"

// myPointer to gives pointer data
func myPointer() {
	var ptr *int = nil
	val := 10

	ptr = &val
	fmt.Println("Address: ", ptr, " Val: ", *ptr)

	ptr1 := &ptr
	fmt.Println("Address: ", *ptr1, " Val: ", **ptr1)
}

// HelloUser simple function for function pointer
func HelloUser(user string) string {
	return "Hello " + user
}

// funPointer to give function Pointer
func funPointer() {
	//var funHello func(string) string = HelloUser
	funHello := HelloUser
	fmt.Println(funHello("World"))
}
