/*
Defer
	- Here defer help in delay of executing perticular function in last
	but before returning function. Any stmt given with defer execute in last.
	- Multiple defer stmt follow LIFO order.
	- It is useful in closing resource at end of function.
	- defer function will take value at that moment so after modification we
	may get lod value only.
		Example:
			a := "start"
			defer fmt.Println(a)
			a = "end"
			----------------------
			Result: start // Get value at that time
	- defer function run before panic statement. It prevent closing resource.
*/

package main

import (
	"fmt"
)

func myDefer1() {
	// start end middle
	fmt.Println("start")
	defer fmt.Println("middle") // execute after end but before function return
	fmt.Println("end")
}

func myDefer2() {
	// LIFO(Reverse): end middle start
	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")
}

func myDefer3() {
	// start as defer point a value is start
	a := "start"
	defer fmt.Println(a)
	a = "end"
}

func myDefer4() {
	// start as defer point a value is start
	defer fmt.Println("start")
	panic("Something Bad Happen !!!")
}

func myDefer() {
	myDefer1()
	fmt.Println()
	myDefer2()
	fmt.Println()
	myDefer3()
	fmt.Println()
	myDefer4()
}

func myControlFlow() {
	myDefer()
}

/*
func main() {
	myControlFlow()
}
*/
