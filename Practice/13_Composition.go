/*
Composition:
	Go dose not suport class inheritance and comosition.
	But using structure we can use composition.

	- We can replicate composition with struct.
	by adding one struct to other.

Example:
	Person {name}
	Student {per Person, roll int}
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
	per   Person // Embedded Person struct in Student created variable like roll
	roll  int
	class string
}

// Teacher type of Person
type Teacher struct {
	per     Person
	subject string
}

func myComposition() {
	stu1 := Student{}
	per := Person{}
	per.name = "A" // Here student is embeding person property
	per.age = 22
	stu1.per = per
	stu1.roll = 2343
	stu1.class = "First Year"

	fmt.Println(stu1)
	fmt.Println(stu1.per.name)
}

func main() {
	myComposition()
}
