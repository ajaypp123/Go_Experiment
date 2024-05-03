package concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
Mutexes (short for mutual exclusion) are used to prevent multiple goroutines from accessing shared resources simultaneously, thus preventing race conditions.

Protecting shared variables in concurrent programs.
Synchronizing access to shared data structures like maps, slices, or other complex objects.
Ensuring data consistency in multi-threaded applications.
*/

type mutexx struct{}

var count int = 0
var mutex sync.Mutex

func (m *mutexx) increment() {
	mutex.Lock()         // Lock the mutex to prevent concurrent access
	defer mutex.Unlock() // Ensure the mutex is unlocked after the function completes
	count += 1
	fmt.Println(count)
}

func (m *mutexx) startPoint() {
	for i := 0; i < 10; i++ {
		go m.increment()
	}
	time.Sleep(time.Second) // Wait for goroutines to finish
}

func TestMutexx(t *testing.T) {
	m := mutexx{}

	m.startPoint()
}
