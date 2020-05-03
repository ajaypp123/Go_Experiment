/*
Channels:
	- Buffered channel used when sender and receiver have different frequency
	- To use different frequency we can use close() and range method, if used close method
	then range will know how many element in buffer and print it.

	- Declear
		ch := make(chan int, 50)
	Here "ch" is channel created to send and receive int data only.

	- Implementation
		i := <- ch // received data from channel
		ch <- i // send data to channel

		// Call function to check if message exist in channel
		go func() {
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		waitGroup.Done()
		}()


*/

package main

import (
	"fmt"
	"strconv"
	"sync"
)

var waitGroup sync.WaitGroup = sync.WaitGroup{}

func senderBuffer(ch2 chan<- string, i string) {
	fmt.Println("Send value: ", i)
	ch2 <- i
	waitGroup.Done()
}

func receiverBuffer(ch2 <-chan string) {
	i := <-ch2
	fmt.Println("Received value: ", i)
	waitGroup.Done()
}

func myBufferedChannel() {
	ch2 := make(chan string, 3)
	for i := 0; i < 5; i++ {
		waitGroup.Add(2)
		go receiverBuffer(ch2)
		go senderBuffer(ch2, strconv.Itoa(i))
	}
	waitGroup.Wait()
}

///////////////////////////////////

func bufferedExample() {
	ch := make(chan int, 50)
	waitGroup.Add(2)

	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		waitGroup.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 11
		ch <- 22
		close(ch)
		waitGroup.Done()
	}(ch)
	waitGroup.Wait()
}

////////////////////////////////////

func bufferedExample1() {
	ch := make(chan int, 50)
	waitGroup.Add(2)
	go func() {
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		waitGroup.Done()
	}()

	go func(ch chan<- int) {
		ch <- 11
		ch <- 22
		close(ch)
		waitGroup.Done()
	}(ch)
	waitGroup.Wait()
}

///////////////////////////////////
/*
func main() {
	myBufferedChannel()
	bufferedExample()
	bufferedExample1()
}
*/
