package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	input := []int{1, 2, 3, 4, 5}
	result := make([]int, len(input))

	for i, data := range input {
		wg.Add(1)
		go proccessData(&wg, &result[i], data)
	}
	wg.Wait()
	fmt.Println(result)
	fmt.Println(time.Since(start))
}

func proccess(data int) int {
	time.Sleep(time.Second)
	return data * 2
}

func proccessData(wg *sync.WaitGroup, resultDestination *int, data int) {
	defer wg.Done()
	processedData := proccess(data)
	*resultDestination = processedData
}
