package main

import (
	"fmt"
	"time"
)

func main() {
	doneChannel := make(chan bool)
	go func() {
		time.Sleep(time.Second * 2)
		doneChannel <- true
	}()

	for {
		select {
		case <-doneChannel:
			fmt.Println("Done")
			return
		default:
			fmt.Println("Still working")
			time.Sleep(time.Second)
		}
	}
}
