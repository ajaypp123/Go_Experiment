package concurrency

import (
	"fmt"
	"testing"
)

/*
State machines are a computational model used to represent and manage the state of a system and the transitions between different states. Using channels in Go, we can implement state machines by communicating state transitions between goroutines.
*/

type stateBlock struct {
	state int //active / completed /
}

// now use
// state := make(chan stateBlock)
// use channel to get state of go routines or process

func TestStateMachine(t *testing.T) {
	stateCh := make(chan stateBlock)
	fmt.Println("1. active, 2. completed")

	// State machine
	go func() {
		fmt.Println("Go routine to check state")
		for {
			select {
			case newState := <-stateCh:
				state := newState
				fmt.Println("New state:", state.state)
			}
		}
	}()

	// Transition to active state
	stateCh <- stateBlock{state: 1}

	// Simulate some work
	fmt.Println("Doing some work...")
	fmt.Println("Completed...")
	// Transition to done state
	stateCh <- stateBlock{state: 2}
}
