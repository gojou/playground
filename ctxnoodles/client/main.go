package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Enter something, dammit.")
	}

	waitseconds, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Enter an integer, dammit.")
	}
	wait := time.Duration(waitseconds * 1000000000)

	fmt.Println(wait)

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	req = req.WithContext(ctx)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}
	io.Copy(os.Stdout, res.Body)
}
