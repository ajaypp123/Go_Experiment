package concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestThrottling(t *testing.T) {
	// Throttle to limit the rate of execution to 2 per second
	throttle := time.Tick(5 * time.Millisecond)

	// Concurrent requests
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			<-throttle // Wait for throttle
			fmt.Printf("Request %d processed\n", id)
		}(i)
	}
	wg.Wait()
}
