package concurrency

import (
	"fmt"
	"testing"
	"time"
)

type buff_channels struct{}

func (c *buff_channels) sendData(ch chan string) {
	for i := 0; i < 5; i++ {
		ch <- fmt.Sprintf("Message %d", i)
		fmt.Println("Send: ", i)
		time.Sleep(100 * time.Millisecond)
	}
	close(ch) // Close the channel after sending all messages
}

func (c *buff_channels) startPoint() {
	ch := make(chan string, 3)

	// Start goroutine to send data
	go c.sendData(ch)

	// Receive data from the channel
	for msg := range ch {
		fmt.Println("Received:", msg)
		time.Sleep(200 * time.Millisecond) // Simulate slower receiver
	}
}

func TestBufferdChannels(t *testing.T) {
	c := buff_channels{}

	c.startPoint()
}
