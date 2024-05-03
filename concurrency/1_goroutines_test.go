package concurrency

import (
	"fmt"
	"testing"
)

type goroutines struct{}

func (g *goroutines) sayHello() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello")
	}
}

func (g *goroutines) startPoint() {
	go g.sayHello()

	go func() {
		fmt.Println("Hi")
	}()

	for i := 0; i < 10; i++ {
		fmt.Println("World")
	}
}

func TestGoroutines(t *testing.T) {
	g := goroutines{}

	g.startPoint()
}
