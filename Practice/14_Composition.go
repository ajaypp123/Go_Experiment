/*
Composition:
Go dose not suport class inheritance and comosition.
But using structure we can use composition.

Example:
	Person {name}
	Student {Person, roll}
	stu.name and stu.roll
Here we are not using person.name and student is embeding person property.
*/

package main

import "fmt"

// Person Base struct
type Person struct {
	name string
	age  int
}

// Student type of Person
type Student struct {
	Person // Embedded Person struct in Student not created variable like roll
	roll   int
	class  string
}

// Teacher type of Person
type Teacher struct {
	Person
	subject string
}

func myComposition() {
	stu1 := Student{}
	stu1.name = "A" // Here student is embeding person property
	stu1.age = 22
	stu1.roll = 2343
	stu1.class = "First Year"

	fmt.Println(stu1)
	fmt.Println(stu1.name)
}

/*
func main() {
	myComposition()
}
*/
