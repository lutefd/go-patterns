package main

import (
	"fmt"
)

func main() {
	msgChannel := make(chan string)
	anotherMsgChannel := make(chan string)
	go func() {
		msgChannel <- "Hello"
	}()
	go func() {
		anotherMsgChannel <- "World"
	}()
	select {
	case msgFromMsgChannel := <-msgChannel:
		fmt.Println(msgFromMsgChannel)
	case msgFromAnotherMsgChannel := <-anotherMsgChannel:
		fmt.Println(msgFromAnotherMsgChannel)
	}

	fmt.Println("Done")
}
