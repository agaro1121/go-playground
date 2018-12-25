package main

import "fmt"

func main() {

	ch := make(chan string, 1)

	ch <- "Saluton Mondo"
	close(ch) // closes sending side of the channel so it can be drained properly

	/*for {
		if msg, isClosed := <- ch; isClosed {
			fmt.Println("msg=",msg)
		} else {
			break
		}
	}*/

	//fmt.Println(<- ch)

	// checks if channel is open and polls message if so
	for msg := range ch {
		fmt.Println("msg2=",msg)
	}

	type Message struct {
		msg string
		code int
	}

	type ErrorMessage struct {
		errorMsg string
		reason string
		code int
	}

	message := Message{
		msg: "some message",
		code: 0,
	}

	errorMessage := ErrorMessage{
		errorMsg: "some error message",
		reason:   "some reason",
		code: -1,
	}

	messageChannel := make(chan Message, 1)
	errorMessageChannel := make(chan ErrorMessage, 1)

	messageChannel <- message
	errorMessageChannel <- errorMessage

	// listen to multiple channels
	// this only gets called once
	// need to put it on a timed loop or something
	select {
		case m := <-messageChannel:
			fmt.Println(m)
		case em := <-errorMessageChannel:
			fmt.Println(em)
		default:
			fmt.Println("default")
		}

}
