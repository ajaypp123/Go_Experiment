package concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
Broadcasting is a synchronization mechanism where a signal is sent to multiple threads or goroutines to wake them up simultaneously. It's commonly used in scenarios where multiple threads need to synchronize their activities.
*/

func TestBroadcast(t *testing.T) {
	var cond sync.Cond
	var wg sync.WaitGroup
	var message string

	cond.L = &sync.Mutex{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		cond.L.Lock()
		defer cond.L.Unlock()

		for message == "" {
			cond.Wait()
		}

		fmt.Println("Listener 1 received:", message)
	}()

	go func() {
		defer wg.Done()
		cond.L.Lock()
		defer cond.L.Unlock()

		for message == "" {
			cond.Wait()
		}

		fmt.Println("Listener 2 received:", message)
	}()

	time.Sleep(1 * time.Second) // Simulate some processing

	cond.L.Lock()
	message = "Hello from broadcaster!"
	cond.Broadcast() // Signal listeners
	cond.L.Unlock()

	wg.Wait()

}
