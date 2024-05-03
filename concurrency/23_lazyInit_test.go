package concurrency

import (
	"fmt"
	"sync"
	"testing"
)

/*
Lazy initialization is a pattern where an object or resource is not created or initialized until it is first accessed or required. This pattern is useful for optimizing resource usage and improving performance by deferring initialization until it is actually needed.
*/

type LazySingleton struct {
	Data string
}

var instance *LazySingleton
var once sync.Once

func getInstance() *LazySingleton {
	once.Do(func() {
		instance = &LazySingleton{Data: "Initialized data"}
	})
	return instance
}

func TestLazyInit(t *testing.T) {
	// Access the singleton instance
	singleton := getInstance()
	fmt.Println("Data:", singleton.Data)
}
