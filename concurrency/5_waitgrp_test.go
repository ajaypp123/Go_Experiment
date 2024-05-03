package concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
Synchronization: WaitGroups are used to wait for a collection of goroutines to finish their execution before proceeding.

Non-Blocking: They don't provide communication capabilities like channels but are purely used for synchronization.

Increment and Decrement: You increment the WaitGroup counter before starting each goroutine, and then decrement it when each goroutine finishes.

Blocking Operation: The Wait method blocks until the WaitGroup counter reaches zero, indicating that all goroutines have finished.
*/

type waitgrp struct{}

func (w *waitgrp) worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("starting task ", i, " ...")
	time.Sleep(2 * time.Second)
	fmt.Println("completed task ", i, " ...")
}

func (w *waitgrp) startPoint() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go w.worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("All workers have finished")
}

func TestWaitGrp(t *testing.T) {
	w := waitgrp{}

	w.startPoint()
}
