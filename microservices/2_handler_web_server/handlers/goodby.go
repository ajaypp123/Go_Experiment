package handlers

import (
	"log"
	"net/http"
)

// GoodBy method
type GoodBy struct {
	l *log.Logger
}

// NewGoodBy method
func NewGoodBy(l *log.Logger) *GoodBy {
	return &GoodBy{l}
}

func (g *GoodBy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Goodby")
	rw.Write([]byte("GoodBy"))
}
