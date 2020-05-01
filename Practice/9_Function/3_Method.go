/*
Method Function:
	- Method are similar to function but it call with known contex. Where known contex
	can be any known type like struct.
	- If method are used for struct then may copying data so best way is to use pointer
	to struct insted normal struct passing.
	Example:
		func (e emp) printData() { // Method that call by employee struct
			// stmt
		}
		e := employee{name:"aaa", age:22}
		e.printData() // Here method can be call with known type
*/

package main

import "fmt"

///////////////////////////////
type employee struct {
	name string
	age  int
}

func (e employee) empData() {
	fmt.Println(e.name, e.age)
}

func (e *employee) empDataPointer() {
	fmt.Println(e.name, e.age)
}

///////////////////////////////

type counter int

func (c counter) incCounter() int {
	c++
	return int(c)
}

func myMethod() {
	// Method for struct
	emp1 := employee{
		name: "aaa",
		age:  22,
	}
	emp1.empData()

	// Method for struct pointer
	emp1.empDataPointer()

	// Method for int
	count := counter(22)
	fmt.Println(count.incCounter())
	fmt.Println(count)
}

func main() {
	myMethod()
}
