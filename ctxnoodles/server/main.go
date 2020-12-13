package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var duration time.Duration

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Enter something, dammit.")
	}

	waitseconds, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Enter an integer, dammit.")
	}
	duration = time.Duration(waitseconds * 1000000000)

	fmt.Println(duration)

	fmt.Println("Server is up...")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {

	defer log.Print("Server has served")

	time.Sleep(duration)
	fmt.Fprintln(w, "Hey")
}
