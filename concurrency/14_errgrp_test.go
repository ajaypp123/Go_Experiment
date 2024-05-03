package concurrency

import (
	"context"
	"fmt"
	"testing"
	"time"

	"golang.org/x/sync/errgroup"
)

/*
The errgroup.Group type provided by the sync/errgroup package allows you to execute multiple functions concurrently as goroutines and wait for all of them to complete. If any of the goroutines return an error, the Wait method of the Group will return the first non-nil error encountered.
*/

type egrp struct{}

func (e *egrp) startPoint() {
	// Create a new errgroup
	g, ctx := errgroup.WithContext(context.Background())

	// Execute multiple functions concurrently
	for i := 0; i < 5; i++ {
		index := i // Capture the loop variable
		g.Go(func() error {
			return e.worker(ctx, index)
		})
	}

	// Wait for all functions to complete
	if err := g.Wait(); err != nil {
		fmt.Printf("Error encountered: %v\n", err)
	} else {
		fmt.Println("All workers completed successfully")
	}
}

func (e *egrp) worker(ctx context.Context, id int) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(time.Duration(id) * time.Second):
		fmt.Printf("Worker %d completed\n", id)
		return nil
	}
}

func TestGerr(t *testing.T) {
	e := egrp{}

	e.startPoint()
}
