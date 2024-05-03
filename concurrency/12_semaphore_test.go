package concurrency

import (
	"fmt"
	"testing"
	"time"
)

type semaphore struct {
	permits chan struct{}
}

func (s *semaphore) aquire() {
	<-s.permits
}

func (s *semaphore) release() {
	s.permits <- struct{}{}
}

func testWorker(sem *semaphore, id int) {
	fmt.Println("Worker", id, "acquiring permit")
	sem.aquire()

	// Simulate work on the resource
	fmt.Println("Worker", id, "working...")
	time.Sleep(time.Second)

	fmt.Println("Worker", id, "releasing permit")
	sem.release()
}

func getSemaphore(n int) *semaphore {
	return &semaphore{permits: make(chan struct{}, n)}
}

func TestSemaphore(t *testing.T) {
	sem := getSemaphore(2) // Allow 2 concurrent workers

	for i := 0; i < 5; i++ {
		go testWorker(sem, i+1)
	}

	time.Sleep(5 * time.Second)
}
