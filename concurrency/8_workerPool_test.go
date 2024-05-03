package concurrency

import (
	"fmt"
	"sync"
	"testing"
)

/*
The Worker Pool pattern is a technique for managing concurrent tasks efficiently. It involves a fixed number of worker goroutines that continuously fetch tasks from a shared queue and execute them. This allows you to process multiple tasks in parallel without creating an excessive number of goroutines, which can lead to performance overhead.

Components:
Workers (Goroutines): These are lightweight threads that represent the independent units of execution within the pool. Each worker defines a function that performs the actual work on the tasks.

Task Queue (Channel): A channel acts as a buffer, holding tasks that need to be processed. Workers continuously receive tasks from the queue using the channel's receive operation (<-chan).
*/

// Task represents the unit of work to be done by a worker
type Task func()

// WorkerPool manages a pool of workers and a task queue
type WorkerPool struct {
	workers   []*worker
	taskQueue chan Task
	wg        sync.WaitGroup
}

// worker represents a single worker goroutine
type worker struct {
	pool *WorkerPool
}

// NewWorkerPool creates a new WorkerPool with a specified number of workers
func NewWorkerPool(numWorkers int) *WorkerPool {
	pool := &WorkerPool{
		workers:   make([]*worker, numWorkers),
		taskQueue: make(chan Task),
	}

	// Create and start worker goroutines
	for i := 0; i < numWorkers; i++ {
		worker := &worker{pool: pool}
		pool.workers[i] = worker
		go worker.run()
	}

	return pool
}

// run is the main loop for each worker goroutine
func (w *worker) run() {
	for {
		task, ok := <-w.pool.taskQueue
		if !ok {
			break // Channel closed, worker exits
		}
		// Execute the task
		task()
		w.pool.wg.Done() // Signal task completion
	}
}

// SubmitTask adds a task to the pool's queue
func (wp *WorkerPool) SubmitTask(task Task) {
	wp.taskQueue <- task
	wp.wg.Add(1) // Track pending tasks
}

// Wait waits for all submitted tasks to finish
func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
	close(wp.taskQueue) // Close the queue after all tasks are submitted
}

func TestWorkerPool(t *testing.T) {
	pool := NewWorkerPool(4) // Create a pool with 4 workers

	// Define tasks
	tasks := []Task{
		func() { fmt.Println("Task 1") },
		func() { fmt.Println("Task 2") },
		func() { fmt.Println("Task 3") },
		func() { fmt.Println("Task 4") },
		// ... add more tasks
	}

	// Submit tasks to the pool
	for _, task := range tasks {
		pool.SubmitTask(task)
	}

	// Wait for all tasks to finish
	pool.Wait()

	fmt.Println("All tasks completed!")
}
