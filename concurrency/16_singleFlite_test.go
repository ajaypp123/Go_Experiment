package concurrency

import (
	"fmt"
	"testing"
	"time"

	"golang.org/x/sync/singleflight"
)

/*
Single flight is a concurrency pattern used to coalesce duplicate requests into a single shared result. It's particularly useful in scenarios where multiple concurrent requests are made for the same expensive operation, such as database queries or API calls.
*/

type sFlight struct{}

func (s *sFlight) startPoint() {
	var sf singleflight.Group

	// Execute the function f with single flight
	v, err, _ := sf.Do("key", func() (interface{}, error) {
		// Simulate an expensive operation
		time.Sleep(1 * time.Second)
		return "result", nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", v)
	}
}

func TestSingleFlight(t *testing.T) {
	s := sFlight{}
	s.startPoint()
}
