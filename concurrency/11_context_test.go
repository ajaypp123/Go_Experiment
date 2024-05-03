package concurrency

import (
	"context"
	"fmt"
	"testing"
	"time"
)

/*
The context package provides a way to propagate deadlines, cancelation signals, and other request-scoped values across API boundaries and between goroutines.

It can intrrupt goroutines.
*/

type contextTest struct{}

func (c *contextTest) worker(_ context.Context, ch chan string) {

	time.Sleep(2 * time.Second)
	ch <- "Task"
}

func (c *contextTest) startPoint() {
	ctx := context.Background()
	ctx, cancle := context.WithTimeout(ctx, time.Second)
	defer cancle()

	ch := make(chan string)
	go c.worker(ctx, ch)

	select {
	case <-ctx.Done():
		fmt.Println("Task cancelled")
	case <-ch:
		fmt.Println("Task Completed")
	}
	fmt.Println("Completed...")
}

func TestContext(t *testing.T) {
	c := contextTest{}

	c.startPoint()
}
