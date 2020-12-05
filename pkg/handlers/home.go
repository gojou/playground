package handlers

import (
	"fmt"
	"net/http"
)

// Home displays the default "/" page
func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Home sweet home.")

}
