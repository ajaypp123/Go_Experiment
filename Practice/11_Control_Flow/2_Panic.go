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
	- Defer function execute just before panic so, it can be used to find panic. but
	still program end with panic. It can be used to print msg like server under maintenance...

Recover
	- Defer function execute just before panic so, it can be used to find panic. but
	still program end with panic. It can be used to print msg like server under maintenance...
	- When panic happen we can get that as error and handle it.
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
	//fmt.Println("end")
}

func myPanic3() {
	/*
		Here end will not print mean application will end.
		But it used to print proper message. like service under maintenance.
	*/
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			//Panic happen need to handle panic
			fmt.Println("Panic happen")
		}
	}()
	a := 0
	b := 20
	fmt.Println(b / a)
	fmt.Println("After Panic")
}

func myPanic() {
	//myPanic1()
	//myPanic2()
	myPanic3()
}

/*
func main() {
	myPanic()
}
*/
