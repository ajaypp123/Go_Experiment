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
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Go_Experiment/microservices/3_rest_service/handlers"
)

func main() {
	l := log.New(os.Stdout, "Hello-pro", log.LstdFlags)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", handlers.NewHello(l))
	serveMux.Handle("/goodby", handlers.NewGoodBy(l))

	server := &http.Server{
		Addr:        "localhost:3000",
		Handler:     serveMux,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
	}

	log.Println("starting server...")

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate sig: ", sig)

	if tc, err := context.WithTimeout(context.Background(),
		30*time.Second); err != nil {
		l.Fatal(err)
	} else {
		server.Shutdown(tc)
	}
	log.Println("server is closed")
}
