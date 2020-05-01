/*
Struct

	- Structure pointer point to struct data
	- In struct a = b will create new object. So passing struct in parameter will
	create new memory.
		Example: testStructPointer3()

Slices
	- In Slices a = b will point to memory of Slices. so changes in a will affect in b.
	- we can use Slices to pass in parameter
		Example: testSlicesPointer1()

Array
	- IN Array a = b will create new object. So passing Array in parameter will create new memory.
		Example: testArrayPointer1()

Map
	- In Map a = b will point to memory of Map. so changes in a will affect in b.
		- we can use Map to pass in parameter
			Example: testMapPointer1()
*/

package main

import "fmt"

type employee struct {
	name string
	age  int
}

func testStructPointer1() {
	//Here we get &{abc 22} where & has its meaning like var pointing to address
	var emp1Ptr *employee
	fmt.Println(emp1Ptr)
	emp1Ptr = &employee{name: "abc", age: 22}
	fmt.Println(emp1Ptr)
}

func testStructPointer2() {
	//Here we get &{abc 22} where & has its meaning like var pointing to address
	var emp1Ptr *employee
	emp1Ptr = new(employee) // Create empty struct and return address
	fmt.Println(emp1Ptr)    // assign default value
	// Go compiler dereference assign value to it
	emp1Ptr.name = "aaa"
	emp1Ptr.age = 22
	fmt.Println(emp1Ptr)
}

func testStructPointer3() {
	fmt.Println("Struct")
	// Here em1 and emp2 will get different memory and emp1 change will not change emp2
	var emp1 employee = employee{name: "abc", age: 22}
	var emp2 employee = emp1
	fmt.Println(emp1, emp2, &emp1, &emp2)
	emp1.name = "fff"
	fmt.Println(emp1, emp2, &emp1, &emp2)
}

func testSlicesPointer1() {
	fmt.Println("Slices")
	a := []int{1, 2, 3}
	b := a
	fmt.Println(a, b, &a, &b)
	b[1] = 22
	b[2] = 33
	fmt.Println(a, b, &a, &b)
}

func testArrayPointer1() {
	fmt.Println("Array")
	var a [3]string
	a[0] = "aaa"
	a[1] = "bbb"
	a[2] = "ccc"
	b := a
	fmt.Println(a, b, &a, &b)
	b[1] = "fff"
	b[2] = "ggg"
	fmt.Println(a, b, &a, &b)
}

func testMapPointer1() {
	fmt.Println("Map")
	a := map[string]int{"aaa": 1, "bbb": 2, "ccc": 3}
	b := a
	fmt.Println(a, b, &a, &b)
	b["aaa"] = 11
	b["ccc"] = 33
	fmt.Println(a, b, &a, &b)
}

func myDataTypePointer() {
	testStructPointer1()
	testStructPointer2()
	testStructPointer3()

	testSlicesPointer1()

	testArrayPointer1()

	testMapPointer1()
}

/*
func main() {
	myDataTypePointer()
}
*/
