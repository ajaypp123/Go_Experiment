// variable, const
/*
Variable: 3 step of visibility
----------------------------------
Global : available accrocess package. must be declear with capital letter.
Package : available inside package. must be declear with small letter outside of function
Block : available inside block/function. must be declear with small letter inside of function

Use of variable case
-----------------------------------
small length (like i, j) variable have shorter lifespam. Longer name (like Employee) mean useing for log time.

Conversion
-----------------------------------
Go will depend on user to covert variable.
Example: var a int8 = 22, var b int32 = 33. a-b need coversion.
a - int8(b)

Constant
------------------------------------
constant operation can be perform
const a = 10; int8 b = 10
a+ b is possible as compiler knows a value
*/

package main

import (
	"fmt"
	"strconv"
)

// I is Global variable
var I int = 11

func declareVariable() {
	/*
		Decdeclare of variable
	*/
	var name string = "ABC"
	var price1, price2 float64 = 22.22, 33.44 //declare variable
	price3 := 30                              // declare variable
	price3 = 22                               // reassign value
	const pi float64 = 3.14                   // const variable

	fmt.Println(name, price1, price2, price3, pi)

	/*
		Set of variable
	*/
	var (
		name1 string
		age   int
	)

	name1 = "abc"
	fmt.Println(name1, age)
}

func conversionVariable() {
	/*
		int to float float to int
	*/
	intA := 44
	floatA := 20.20
	var floatB float64 = float64(intA)
	var intB int = int(floatA)
	fmt.Println(intA, floatA, intB, floatB)

	/*
		strint to int and int to string.
		if we use string() it will replace num with unicode char.
		Use strcov lib for string coversion
	*/
	var strA string = string(intA)
	var strB string = strconv.Itoa(intB)
	var strC string = "222"
	intB, _ = strconv.Atoi(strC)
	fmt.Println(strA, strB, strC, intB)
}

func primitiveVariable() {
	/*
		bool: default is false.
		signed int: int8, int16 ... int64
		unsigned int: uint16
	*/
	var x uint16 = 11
	fmt.Println(x)

	/*
		Shift operator.
	*/
	var y int32 = 27
	fmt.Println(y<<3, y>>3)
}

func constVariable() {
	/*
		itoa will increase value from 0 to 1 by 1
	*/
	const (
		a = iota
		b
		c
	)
	fmt.Println(a, b, c)
}

func myVariable() {
	declareVariable()
	conversionVariable()
	primitiveVariable()
	constVariable()
}

/*
func main() {
	myVariable()
}
*/
