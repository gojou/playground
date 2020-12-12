package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	sleepAndTalk(ctx, 4*time.Second, "Hola!")

}

func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}
