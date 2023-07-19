package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan int)

	go func() {
		messages <- 123
	}()

	// wait

	fmt.Println(<-messages)

	bufferedMessages := make(chan int, 5)
	bufferedMessages <- 123
	bufferedMessages <- 234
	bufferedMessages <- 345

	fmt.Println(<-bufferedMessages)
	fmt.Println(<-bufferedMessages)
	time.Sleep(time.Second)
	fmt.Println(<-bufferedMessages)
}
