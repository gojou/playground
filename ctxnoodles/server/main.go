package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Server listening")
	defer log.Print("Server has served")

	time.Sleep(time.Second * 2)
	fmt.Fprintln(w, "Hey")
}
