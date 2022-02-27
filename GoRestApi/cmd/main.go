package main

import (
	"log"
	"net/http"

	"github.com/ajaypp123/GoRestApi/controller"
	"github.com/gorilla/mux"
)

func initalizeRouter() {
	r := mux.NewRouter()

	controller.InitUser()
	r.HandleFunc("/users", controller.GetUsers).Methods("GET")
	r.HandleFunc("/user/{id}", controller.GetUser).Methods("GET")
	r.HandleFunc("/user", controller.PutUser).Methods("POST")
	r.HandleFunc("/user/{id}", controller.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", controller.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	initalizeRouter()
}
