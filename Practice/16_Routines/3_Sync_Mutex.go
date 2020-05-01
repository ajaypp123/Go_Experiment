/*
Mutex
	- It is Lock to application variable that shared among multiple routine.
	- WaitGroup is simple with small scale but if there are 30 to 40 routine sharing
	variable among them then it can cause race condition.
	- variable don't have lock and can be access by any routine radomely.
	- To avoid these we can use mutex variable locking.

	- Type of Mutex:
	1. Mutex:
		- It is common mutex mean at time only one routine are allowed to access
		it and other has to wait for their chance.
	2. RWMutex:
		- Anyone can read this data but only one can write it at a time.
		- Once write operation come it will wait till all reader in queue to finish and
		then writer will held lock and reader will wait till lock release.
	- Mutex lock must be lock before starting routine and release inside of routine
	Example:
		func fun(){
			mutex.Lock()
			go fun2() // Unlock mutex inside of fun2 after work done
		}

		func fun2(){
			//do something with shared variable
			mutex.Unlock()
			// do something with local variable
		}

	Declear:
		var rwMutex sync.RWMutex = sync.RWMutex{}
		var mutex sync.Mutex = sync.Mutex{}

		mutex.Lock()
		mutex.Unlock()

		rwMutex.Lock()
		rwMutex.Unlock()
		rwMutex.RLock()
		rwMutex.RUnlock()
*/

package main

import (
	"fmt"
	"sync"
)

var counter int = 0
var waitGrp sync.WaitGroup = sync.WaitGroup{}

////////////////////////////////////
func incWG() {
	counter++
	waitGrp.Done()
}

func helloWGRoutine() {
	fmt.Println("Hello Counter is : ", counter)
	waitGrp.Done()
}

func myWaitGroupRoutine() {
	for i := 0; i < 10; i++ {
		waitGrp.Add(2)
		go incWG()
		go helloWGRoutine()
	}
	waitGrp.Wait()
}

///////////////////////////////////////

var rwMutex sync.RWMutex = sync.RWMutex{}

func incWG1() {
	counter++
	rwMutex.Unlock()
	waitGrp.Done()
}

func helloWGRoutine1() {
	fmt.Println("Hello Counter is : ", counter)
	rwMutex.RUnlock()
	waitGrp.Done()
}

func myWaitGroupRoutine1() {
	for i := 0; i < 10; i++ {
		waitGrp.Add(2)
		rwMutex.Lock()
		go incWG1()
		rwMutex.RLock()
		go helloWGRoutine1()
	}
	waitGrp.Wait()
}

/*
func main() {
	// runtime.GOMAXPROCS(1) // Set to 1 will run application as single threaded.
	myWaitGroupRoutine()
	fmt.Println()
	myWaitGroupRoutine1()
}
*/
