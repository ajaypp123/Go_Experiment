/*
Pointer for struct

	- Structure pointer point to struct data
	- New keyward used to create memory for datatype in heap storage.
	- Data created with new will have there own default value
*/

package main

import "fmt"

type employee1 struct {
	name string
	age  int
}

func testNewPointer() {
	//Here we get &{abc 22} where & has its meaning like var pointing to address
	var emp1Ptr *employee1
	emp1Ptr = new(employee1) // Create empty struct and return address
	fmt.Println(emp1Ptr)     // assign default value
	// Go compiler dereference assign value to it
	emp1Ptr.name = "aaa"
	emp1Ptr.age = 22
	fmt.Println(emp1Ptr)
}

func myNewPointer() {
	testNewPointer()
}

/*
func main() {
	myNewPointer()
}
*/
