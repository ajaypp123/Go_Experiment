/*
Listen to server:
	http.ListenAndServe("localhost:3000", handlers)

Handle request:
	http.HandleFunc

Request
	to read all parameter and data comming from request

*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		if data, err := ioutil.ReadAll(r.Body); err != nil {
			log.Println(err)
		} else {
			log.Printf("%s", data)
			fmt.Fprintf(rw, "Responce: %s\n", data)
		}
	})

	http.HandleFunc("/goodby", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Goodby")
	})

	log.Println("starting server...")
	http.ListenAndServe("localhost:3000", nil) // handlers
	log.Println("server is closed")
}
