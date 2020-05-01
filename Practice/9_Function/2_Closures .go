/*
Closures Function:
	- The function which return Anonymous function pointer.
	Here after getting pointer it keep track of its call.
	- Outer inner function can access outer function variable but reverse is not true.
	- In Async call accessing outer function parameter directly may not good as call async
	may change value.
	- Once pointer created then it keep track of its variable.
		Example:
			func incrementClosure() func()int{
				i := 0
				return func int() {
					i++
					return i
				}
			}
			incrementFun = incrementClosure()
			incrementFun() //1
			incrementFun() //2

*/

package main

import "fmt"

// Anonymous function
func Anonymous() {
	func(msg string) {
		fmt.Println(msg)
	}("In Anonymous Function")
}

// Closures function
func Closures() func(string) {
	return func(msg string) {
		fmt.Println(msg)
	}
}

// incrementClosures keep track of Anonymous function and increment value
func incrementClosures() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// myClosures to call all function
func myClosures() {
	Anonymous()
	funClosures := Closures()
	funClosures("In Closures Function")

	// Increment Function
	funIncClosures := incrementClosures()
	fmt.Println(funIncClosures())
	fmt.Println(funIncClosures())
	fmt.Println(funIncClosures())
	fmt.Println(funIncClosures())
	fmt.Println(funIncClosures())
}
