package main

import "fmt"

func main() {
	charCannel := make(chan string, 3)
	chars := []string{"a", "b", "c"}
	for _, char := range chars {
		select {
		case charCannel <- char:
		default:
			fmt.Println("Channel is full")
		}
	}
	close(charCannel)

	for result := range charCannel {
		fmt.Println(result)
	}
}
