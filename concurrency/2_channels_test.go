package concurrency

import (
	"fmt"
	"testing"
	"time"
)

/*
* Communication: Channels are primarily used for communication and synchronization between goroutines.

* Blocking Operations: Sending or receiving from a channel can block until the other party is ready, allowing for synchronization.

* Buffering: Channels can be buffered, allowing for a fixed number of elements to be stored before blocking.
 */

type channels struct{}

func (c *channels) sendData(ch chan string) {
	for i := 0; i < 5; i++ {
		ch <- fmt.Sprintf("Message %d", i)
		fmt.Println("Send: ", i)
		time.Sleep(100 * time.Millisecond)
	}
	close(ch) // Close the channel after sending all messages
}

func (c *channels) startPoint() {
	ch := make(chan string)

	// Start goroutine to send data
	go c.sendData(ch)

	// Receive data from the channel
	for msg := range ch {
		fmt.Println("Received:", msg)
		time.Sleep(200 * time.Millisecond) // Simulate slower receiver
	}
}

func TestChannels(t *testing.T) {
	c := channels{}

	c.startPoint()
}
