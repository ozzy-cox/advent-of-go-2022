package main

import (
	"fmt"
	"time"
)

func main() {
	bufferedMessages := make(chan int, 1)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("done")
		time.Sleep(time.Second / 2)

		bufferedMessages <- 1
	}()
	<-bufferedMessages
	fmt.Println("dunner")
}
