package main

import (
	"fmt"
	"time"
)

func myHello() {
	//fmt.Scanln()
	fmt.Println("Before time: ", time.Now())
	fmt.Println("Hello World")
	fmt.Println("After time", time.Now())
	fmt.Println(4 == 4, 4 == 3)
}
