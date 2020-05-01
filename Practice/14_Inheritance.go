/*
Inheritance:
	Go dose not suport class inheritance directly.
	But using structure we can use inheritance.

	- In structure we can use other structure ref and
	go will inherit one structure property to other.

Example:
	Person {name}
	Student {Person, roll}
	stu.name and stu.roll
Here we are not using person.name and student is embeding person property.
*/

package main

import "fmt"

// Person Base struct
type person struct {
	name string
	age  int
}

// Student type of Person
type student struct {
	Person // Embedded Person struct in Student not created variable like roll
	roll   int
	class  string
}

// Teacher type of Person
type teacher struct {
	Person
	subject string
}

func myInheritance() {
	stu1 := student{}
	stu1.name = "A" // Here student is inheriting person property
	stu1.age = 22
	stu1.roll = 2343
	stu1.class = "First Year"

	fmt.Println(stu1)
	fmt.Println(stu1.name)
}

/*
func main() {
	myInheritance()
}
*/
