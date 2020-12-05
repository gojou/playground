package main

import (
	"github.com/gojou/bones/pkg/handlers"
	"github.com/gorilla/mux"
)

func routes(r *mux.Router) {
	r.HandleFunc("/", handlers.Home)
	r.HandleFunc("/about", handlers.About)
	//	r.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

}
