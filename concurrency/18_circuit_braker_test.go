package concurrency

import (
	"fmt"
	"testing"
	"time"

	"github.com/sony/gobreaker"
)

func TestCircuitBraker(t *testing.T) {
	// Create a circuit breaker
	cb := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "example",
		MaxRequests: 3,
		Interval:    5 * time.Second,
	})

	// Simulate calls to a failing service
	for i := 0; i < 5; i++ {
		result, err := cb.Execute(func() (interface{}, error) {
			return simulateServiceCall(i)
		})
		if err != nil {
			fmt.Printf("Error: Circuit breaker is open, service call failed: %v\n", err)
		} else {
			fmt.Printf("Success: Service call result: %v\n", result)
		}
		time.Sleep(2 * time.Second) // Simulate some delay between service calls
	}
}

// simulateServiceCall simulates a service call
func simulateServiceCall(i int) (string, error) {
	if i%2 == 0 {
		return fmt.Sprintf("Result %d", i), nil // Success
	}
	return "", fmt.Errorf("Service error") // Simulate a service error
}
