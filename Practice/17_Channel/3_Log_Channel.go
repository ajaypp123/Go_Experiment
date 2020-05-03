/*
// struct with no data is just signal to close application and it is zero memory allocation data.
var endCh = make(chan struct{})
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var endCh = make(chan struct{})
var wait = sync.WaitGroup{}

func logger() {
	for {
		select {
		case entry := <-logCh:
			fmt.Println(entry.time, " [", entry.severity, "]: ", entry.message)
		case <-endCh:
			break
		}
	}
}

func myApplication() {
	go logger()
	logCh <- logEntry{time.Now(), logInfo, "This is Info"}
	logCh <- logEntry{time.Now(), logInfo, "This is Info"}
	logCh <- logEntry{time.Now(), logWarning, "This is Warn"}
	time.Sleep(100000)
	endCh <- struct{}{}
}

func main() {
	myApplication()
}
