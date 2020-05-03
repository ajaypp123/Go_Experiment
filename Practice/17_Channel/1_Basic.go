/*
Channels:
	- Many language come before concurrency and parallism so they have
	library to handle this sync and thread.
	- Go invented in multi-core world so, it has in-build concurrency and parallism library

	- Channel help to sync data between multiple go routine and handling race condition.
	- Channel is like pipe between sender and receiver.
	- If we send five message and receive one message only then it can create deadlock as
	we deal with unbuffered channel
	- Buffered channel used when sender and receiver have different frequency

	- Declear
		ch := make(chan int) // Unbuffered channel
		ch := make(chan int, 50)
	Here "ch" is channel created to send and receive int data only.

	- Implementation
		i := <- ch // received data from channel
		ch <- i // send data to channel

	- Passing channel to function
		chan <-          writing to channel (output channel)
		<- chan          reading from channel (input channel)
		chan             read from or write to channel (input/output channel)
*/

package main

import (
	"fmt"
	"strconv"
	"sync"
)

var wg sync.WaitGroup = sync.WaitGroup{}

/////////////////////////////////////////
func sender(ch chan string, i string) {
	fmt.Println("Send value: ", i)
	ch <- i
	status := <-ch
	fmt.Println("status: ", status)
	wg.Done()
}

func receiver(ch chan string) {
	i := <-ch
	fmt.Println("Received value: ", i)
	ch <- "OK for " + i
	wg.Done()
}

func myBidirectChannel() {
	ch := make(chan string)
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go receiver(ch)
		go sender(ch, strconv.Itoa(i))
	}
	wg.Wait()
}

//////////////////////////////////////

func senderOnly(ch1 chan<- string, i string) {
	fmt.Println("Send value: ", i)
	ch1 <- i
	wg.Done()
}

func receiverOnly(ch1 <-chan string) {
	i := <-ch1
	fmt.Println("Received value: ", i)
	wg.Done()
}

func mySingleDirectChannel() {
	ch1 := make(chan string)
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go receiverOnly(ch1)
		go senderOnly(ch1, strconv.Itoa(i))
	}
	wg.Wait()
}

///////////////////////////////////
/*
func main() {
	myBidirectChannel()
	mySingleDirectChannel()
}
*/
