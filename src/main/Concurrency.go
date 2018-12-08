package main

// goroutines
// Actor Model
// goroutines = Actor
// actors communicate via channels

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	// runs on separate threads/cores
	// probably don't need this for
	// this example
	runtime.GOMAXPROCS(2)

	// block main thread
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	// adding `go` makes it a goroutine
	go func() {
		defer waitGrp.Done()

		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()

	// adding `go` makes it a goroutine
	go func() {
		defer waitGrp.Done()

		fmt.Println("beef")
	}()

	waitGrp.Wait()

	// channels
	/*
		This creates a channel that
		2 goroutines can communicate on
		The 1st goroutine will put messages on the channel
		it will wait (block) until another goroutines
		picks up this message. This is effectively
		synchronous behavior
	*/
	unbufferedChannel := make(chan int)
	fmt.Println(unbufferedChannel)

	/*
		More async
		If channel is full, then goroutine
		will block until it can put message
		on channel
	*/
	bufferedChannel := make(chan int, 5)
	fmt.Println(bufferedChannel)

	/*
		on the receiving side,
		if channel is empty, goroutine will block
		until it picks up a message from channel
		this behavior is same for buffered and unbuffered
	*/
}
