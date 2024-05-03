package concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// BlockingQueue is a thread-safe queue with blocking operations
type BlockingQueue[T any] struct {
	mutex    sync.Mutex
	cond     sync.Cond
	queue    []T
	capacity int
}

// NewBlockingQueue creates a new BlockingQueue with a specified capacity
func NewBlockingQueue[T any](capacity int) *BlockingQueue[T] {
	return &BlockingQueue[T]{
		mutex:    sync.Mutex{},
		cond:     sync.Cond{L: &mutex},
		queue:    make([]T, 0, capacity),
		capacity: capacity,
	}
}

// Push adds an element to the queue. It blocks if the queue is full.
func (q *BlockingQueue[T]) Push(data T) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	for len(q.queue) == q.capacity {
		q.cond.Wait()
	}
	q.queue = append(q.queue, data)
	q.cond.Signal() // Signal waiting consumers
}

// Pop removes and returns an element from the queue. It blocks if the queue is empty.
func (q *BlockingQueue[T]) Pop() T {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	for len(q.queue) == 0 {
		q.cond.Wait()
	}
	data := q.queue[0]
	q.queue = q.queue[1:]
	q.cond.Signal() // Signal waiting producers (optional)
	return data
}

func TestBlockingQueue(t *testing.T) {
	queue := NewBlockingQueue[string](2) // Queue with capacity 2

	// Producer goroutine
	go func() {
		for i := 0; i < 5; i++ {
			data := fmt.Sprintf("Data %d", i)
			queue.Push(data)
			fmt.Println("Producer added:", data)
			time.Sleep(time.Second)
		}
	}()

	// Consumer goroutine
	go func() {
		for i := 0; i < 5; i++ {
			data := queue.Pop()
			fmt.Println("Consumer received:", data)
			time.Sleep(2 * time.Second)
		}
	}()

	time.Sleep(10 * time.Second)
}
