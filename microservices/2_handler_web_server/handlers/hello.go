package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello to handle log
type Hello struct {
	l *log.Logger
}

// NewHello to provide log struct
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if data, err := ioutil.ReadAll(r.Body); err != nil {
		h.l.Println(err)
		http.Error(rw, "Oops", http.StatusBadRequest)
	} else {
		h.l.Printf("%s", data)
		fmt.Fprintf(rw, "Responce: %s\n", data)
	}
}
