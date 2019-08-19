package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup // 1

func main() {
	channel := make(chan int,1000)
	for i := 0; i < cap(channel); i++ {
		wg.Add(1)
		go printAndSleep(i, channel)
	}
	wg.Wait()
	close(channel)
	for number := range channel{
		fmt.Print(number,",")
	}

}

func printAndSleep(number int, channel chan int) {
	defer wg.Done()
	sleepTime := rand.Intn(3)
	time.Sleep(time.Duration(sleepTime)* time.Second)
	channel <- number
}
