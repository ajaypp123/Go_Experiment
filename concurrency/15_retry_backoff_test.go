package concurrency

import (
	"fmt"
	"testing"
	"time"
)

/*
Retry with backoff is a technique used to handle transient errors by automatically retrying failed operations with increasing delays between retries. Backoff strategies typically involve exponentially increasing the delay between retries to avoid overwhelming the system with retries if the underlying issue persists.
*/

func retryWithBackoff(fn func() error, maxAttempts int,
	baseDelay time.Duration) error {
	for attempt := 1; ; attempt++ {
		if err := fn(); err == nil {
			return nil
		} else if attempt >= maxAttempts {
			return fmt.Errorf("failed after %d attempts: %v", maxAttempts, err)
		}

		delay := baseDelay * time.Duration(attempt)
		fmt.Printf("Attempt %d failed. Retrying after %v...\n", attempt, delay)
		time.Sleep(delay)
	}
}

func TestBackOffRetry(t *testing.T) {
	// Example usage
	err := retryWithBackoff(func() error {
		// Simulate a failed operation
		return fmt.Errorf("transient error")
	}, 5, time.Second)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Operation succeeded")
	}
}
