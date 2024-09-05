package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	input := []int{1, 2, 3, 4, 5}
	result := []int{}

	for _, data := range input {
		wg.Add(1)
		go proccessData(&wg, &result, data)
	}
	wg.Wait()
	fmt.Println(result)
	fmt.Println(time.Since(start))

}
func proccess(data int) int {
	time.Sleep(time.Second)
	return data * 2
}
func proccessData(wg *sync.WaitGroup, result *[]int, data int) {
	defer wg.Done()
	processedData := proccess(data)
	lock.Lock()
	*result = append(*result, processedData)
	lock.Unlock()
}
