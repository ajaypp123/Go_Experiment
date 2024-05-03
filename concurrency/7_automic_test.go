package concurrency

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

/*
Atomic operations ensure that a specific operation on a shared variable is executed atomically, without interruption. This prevents data races and ensures data integrity in concurrent programs.
*/

type atomicInt struct {
	c1 int64
	c2 int64
}

func (a *atomicInt) incrementAutomic() {
	atomic.AddInt64(&a.c1, 1)
}

func (a *atomicInt) incrementNonAutomic() {
	a.c2 += 1
}

func (a *atomicInt) worker(wg *sync.WaitGroup) {
	defer wg.Done()
	go a.incrementAutomic()
	go a.incrementNonAutomic()
}

func (a *atomicInt) startPoint() {
	var wg sync.WaitGroup
	wg.Add(200)

	for i := 0; i < 100; i++ {
		go a.worker(&wg)
		go a.worker(&wg)
	}

	wg.Wait()
	time.Sleep(time.Second)
	fmt.Println("Final: ", a.c1)
	fmt.Println("Final: ", a.c2)
}

func TestAutomic(t *testing.T) {
	a := atomicInt{
		c1: 0,
		c2: 0,
	}
	a.startPoint()
}
