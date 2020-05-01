/*
Routines:
	- It help to execute concurrent and parallel programing.
	- In parallel programing it provide synchronization.
	- Avoid go routine in library and allow user to handle sync

	- Synchronization
		WaitGroup
		Mutex
	- Parallel
		runtime GOMAXPROCS more than 1

runtime
	- This is go library which help in controlling number of process,
	processor, read trace and much more.
	- runtime.GOMAXPROCS(int) if int is negative then it gives number of core available in
	system.
		fmt.Println(runtime.GOMAXPROCS(-1)) // will print 4 as no of core is 4
	- Positive value will set number of thread that we can work with.
	- Set value to 1 then we will run application as single threaded.


Race Condition
	- When multiple routine accessing shared data. Then may get unpredictable result.
	This condition called Race condition.
	- To find race condition we can use
		"go run --race program.go"

Thread
	Basic
		- Many programming language follow os thread to create separate stack
		and execute it in parallel.
		- Os thread and creation, destruction operation are expensive and large,
		to manage many lang support thread pull.

	Working
		- Go follow different module like green thread to avoid creating expensive thread.
		- Go create abstraction of thread  called go routines.
		- Go has schedular that going to map go routines to os thread as 4 period of time.
		Then schedular will take turn with every cpu thread available and assign go routines to thread
		with some processing time.
		- Benifit of schedular is that it can work with small stack. So running thousand of go routines
		is possible than thousand of thread.

	Implementation
		- using "go funCall()" will create go routines and processed further without knowing routine is
		finished.
		- routine can be executed if put some sleep time after routine.
*/

package main

import (
	"fmt"
	"time"
)

func helloRoutine() {
	// Hello World may not print as go not wait to finish routine
	go fmt.Println("Hello World")
}

func helloRoutine1() {
	// Hello World will print as go wait to finish sleep time give chance to routine
	go fmt.Println("Hello World")
	time.Sleep(10 * time.Millisecond)
}

func myRoutine() {
	helloRoutine1()
	helloRoutine()
}

/*
func main() {
	fmt.Println(runtime.GOMAXPROCS(-1))
	runtime.GOMAXPROCS(100)
	fmt.Println(runtime.GOMAXPROCS(-1))
	myRoutine()
}
*/
