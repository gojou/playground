package routing

import (
	"github.com/gojou/playground/pkg/handlers"
	"github.com/gorilla/mux"
)

// Routes does its thing
func Routes(r *mux.Router) {
	r.HandleFunc("/", handlers.Home)
	r.HandleFunc("/about", handlers.About)
	r.HandleFunc("/scouts", handlers.Scouts)

}
