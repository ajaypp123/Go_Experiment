package concurrency

import (
	"fmt"
	"testing"
	"time"
)

/*
The select statement in Go allows you to wait on multiple communication operations simultaneously. It's useful for coordinating between multiple channels.
*/

type selectChannels struct{}

func (s *selectChannels) worker1(ch chan string) {
	ch <- "Message from worker1"
}

func (s *selectChannels) worker2(ch chan string) {
	ch <- "Message from worker2"
}

func (s *selectChannels) startPoint() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go s.worker1(ch1)
	go s.worker2(ch2)

	for i := 0; i < 2; i++ {
		select {
		case res := <-ch1:
			fmt.Println("Received from channel: ", res)
		case res := <-ch2:
			fmt.Println("Received from channe2: ", res)
		}
	}
	time.Sleep(100 * time.Millisecond)
}

func TestSelectChannel(t *testing.T) {
	s := selectChannels{}

	s.startPoint()
}
