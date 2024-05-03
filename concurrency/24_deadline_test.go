package concurrency

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestDeadline(t *testing.T) {
	// Create a parent context with a deadline
	parentCtx, cancel := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
	defer cancel()

	// Perform concurrent operations with deadlines
	for i := 0; i < 3; i++ {
		go func(ctx context.Context, id int) {
			select {
			case <-ctx.Done():
				fmt.Printf("Goroutine %d: Deadline exceeded\n", id)
			case <-time.After(time.Duration(id) * time.Second):
				fmt.Printf("Goroutine %d: Operation completed\n", id)
			}
		}(parentCtx, i)
	}

	// Wait for some time to allow goroutines to complete
	time.Sleep(5 * time.Second)
}
