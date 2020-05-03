/*
Routine in closures
	- Go routine in closure may create race condition.

	Problem:
		- Example
			func routineClosure() {
				msg := "Hello"
				go func() { fmt.Println(msg) }()
				msg = "Buy"
				time.Sleep(1000)
			}()
		- Here as we can see that msg set to "Hello" and inner function running as
		routine then changing value to "Buy".
		- schedular may run "Buy" or routine randomly create race condition.

	Solution:
		1. Pass value as parameter
			- Insted of accessing it directly from outside. pass it as parameter.
			It will create copy of variable and gives correct output.
		2. Use Waitgroup
			- Add(1) to sync and use Done to remove and use Wait insted of sleep.

WaitGroup
	- It help to avoid race condition in go routine.
	- WaitGroup help schedular to sync multiple go routine together.
	- WaitGroup will start sync from 0, if we want to sync routine then
	we can add wg.Add(1) and at end end of Routine we can add wg.Done()
	- Add will add one more sync to schedular and Done will subtract from schedular.
	- Here Wait() method will wait till count become zero and then go further else stop
	program for finish routine.

	- WaitGroup is simple with small scale but if there are 30 to 40 routine sharing
	variable among them then it can cause race condition.
	- variable don't have lock and can be access by any routine radomely.
	- To avoid these we can use mutex variable locking.

	Example:
		var wg sync.WaitGroup = sync.WaitGroup{}
		wg.Add(1) // Increase sync count to 1
		wg.Done() // Decrease sync count to 0
		wg.Wait() // Wait till count become zero

*/

package main

import (
	"fmt"
	"sync"
	"time"
)

/*
- Here routineClosure can create race condition.
Output:
  it can be "Hello" or "Buy" may change.
*/
func routineClosure1() {
	msg := "Hello"
	go func() {
		fmt.Println(msg)
	}()
	msg = "Buy"
	time.Sleep(1000)
}

/*
Here output always be "Hello" as we are passing cop of variable
insted of ref.
*/
func routineClosure2() {
	msg := "Hello"
	go func(msg string) {
		fmt.Println(msg)
	}(msg)
	msg = "Buy"
	time.Sleep(1000)
}

/*
Use on WaitGroup.
It sync main and closure go routine
*/
var wg sync.WaitGroup = sync.WaitGroup{}

func routineClosure3() {
	msg := "Hello"
	wg.Add(1)
	go func() {
		fmt.Println(msg)
		wg.Done()
	}()
	wg.Wait()
	msg = "Buy"
}

func myClosureRoutine() {
	routineClosure1()
	routineClosure2()
	routineClosure3()
}

func main() {
	myClosureRoutine()
}
