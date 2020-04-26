// Condition : if-else, Switch case
// switch case has no break case
// No ternary operator
/*
if-else has two syntax
---------------------------------------
1.
if condition {
	stmt1
}

2.
if initialize; condition {
	statement
}
*/

package main

import "fmt"

func myCondition() {
	/*
		Condition Example
	*/
	i := 44
	if i <= 30 {
		fmt.Println("less than 30")
	} else if i <= 50 {
		fmt.Println("less than 50")
	} else {
		fmt.Println("more than 50")
	}

	/*
		Indianization Example
	*/
	myMap := map[string]int{"a": 1, "b": 2}
	if c, ok := myMap["c"]; ok {
		fmt.Println("C is present in map ", c)
	} else {
		fmt.Println("C is not present in map", c)
	}
}

func mySwitch() {
	i := 1
	switch i {
	case 1:
		fmt.Println(i)
	case 3, 4, 5:
		fmt.Println(i)
	default:
		fmt.Println(i)
	}

	switch {
	case i < 44:
		fmt.Println("<44", i)
	case i == 1:
		fmt.Println("==1", i)
	}
}

/*
func main() {
	myCondition()
	mySwitch()
}
*/
