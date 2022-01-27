/*
Go anylising tool
	go get -u -v github.com/nsf/gocode
	go get -u -v github.com/rogpeppe/godef
	go get -u -v github.com/golang/lint/golint
	go get -u -v github.com/lukehoban/go-outline
	go get -u -v sourcegraph.com/sqs/goreturns
	go get -u -v golang.org/x/tools/cmd/gorename
	go get -u -v github.com/tpng/gopkgs
	go get -u -v github.com/newhook/go-symbols
	go get -u -v golang.org/x/tools/cmd/guru
*/
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

/*
func main() {
	myHello()
}
*/
