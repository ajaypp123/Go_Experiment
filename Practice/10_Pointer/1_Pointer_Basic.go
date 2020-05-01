/*
Pointer
	- Variable that store address of other variable
	- Default or undeclared pointer point to nil
	- Go by default dose not support pointer arithmetic operation as it can get
	complex and go is simplicity language.
	- If we want to perform pointer operation then we can use "unsafe" package but then
	Go will not manage runtime memory.
	- Example: &a + 4 not supported
*/

package main

import "fmt"

// myPointer to gives pointer data
func myPointer() {
	var ptr *int
	fmt.Println(ptr)
	val := 10

	ptr = &val
	var ptr2 *int = &val
	fmt.Println("Address: ", ptr, " Val: ", *ptr, &ptr, *ptr2)

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
	myPointer()
	fmt.Println(funHello("World"))
}

func main() {
	funPointer()
}
