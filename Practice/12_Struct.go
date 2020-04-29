/*
Struct
It is same as c.
If we want to expose struct outside use start_capital letter for variables.

Copy
--------------------------
Map and slices follow same ref during copy
but Map follow different ref during copy and modified one var not affect other.
If you are passing it as parameter then you are creating other object.
*/
package main

import "fmt"

//Employee Struct
type Employee struct {
	fName, lName string
	id           int
}

func myEmpStruct() {
	emp := Employee{fName: "Aj", lName: "Pa", id: 22}
	emp1 := Employee{"Aj", "Pa", 22}
	fmt.Println(emp, emp1)
	fmt.Println(emp.fName, emp1.fName)

	emp2 := &emp
	fmt.Println(emp2.fName)
}

/*****************************************************/

//Rect struct for rectangle
type Rect struct {
	length int
	width  int
}

func areaRect(r Rect) int {
	return r.length * r.width
}

func parameterRect(r *Rect) int {
	return 2 * (r.length + r.width)
}

func myRectStruct() {
	rect := Rect{22, 11}
	rect1 := &rect
	fmt.Println("Area: ", areaRect(rect))
	fmt.Println("Parameter: ", parameterRect(rect1))
}

/*
func main() {
	myEmpStruct()
	myRectStruct()
}
*/
