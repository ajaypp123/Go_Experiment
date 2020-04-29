/*
Panic
	- Go not have exception like other language. Because lot of exception
	in other language consider as normal in Go language.
	- Example if we open file which not exist it return err message but not create Exception.
	- But in some case Go will not handle error and create panic.
	- There are two type of panic
		1. panic created by program like divided by 0
		2. User defined panic like panic() method
	- Panic cause shutdown program.
	- Default panic are fatal but we can recover it.

Recover
*/

package main

import "fmt"

func myPanic1() {
	a := 0
	b := 20
	fmt.Println(b / a)
}

func myPanic2() {
	fmt.Println("start")
	defer fmt.Println("middle")
	panic("Panic happes !!!")
	fmt.Println("end")
}

func myPanic() {
	//myPanic1()
	myPanic2()
}

func main() {
	myPanic()
}
