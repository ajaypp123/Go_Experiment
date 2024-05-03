package concurrency

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type leaderElect struct {
	ch chan int
}

func (e *leaderElect) election() {
	for i := 0; i < 5; i++ {
		go func(id int) {
			select {
			case <-time.After(time.Duration(rand.Intn(100)) * time.Millisecond):
				e.ch <- id
			}
		}(i)
	}
}

func (e *leaderElect) startPoint() {
	e.election()

	leader := <-e.ch
	fmt.Println(leader)
}

func TestLeaderElection(t *testing.T) {
	e := leaderElect{
		ch: make(chan int),
	}

	e.startPoint()
}
