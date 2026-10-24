package main

import "time"

func main() {
	oneCh := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		oneCh <- "Hello, World!"
	}()

	for i := 0; i < 5; i++ {
		println("Doing some work...")
		println("index: ", i)
	}

	println("Waiting for message...")
	msg := <-oneCh
	println(msg)
}
