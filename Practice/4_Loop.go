// Only for loop in go

package main

import (
	"fmt"
)

func myLoop() {
	i := 0

	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	for k, l := 0, 0; l < 10; l++ {
		fmt.Println(k, l)
	}
}
