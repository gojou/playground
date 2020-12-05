package handlers

import (
	"fmt"
	"net/http"
)

// About displays information about the site
func About(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "About that...")
}
