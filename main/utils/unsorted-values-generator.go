package utils

import (
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup // 1

func GenerateUnsortedValues(size int) (values []int) {
	channel := make(chan int,size)
	for i := 0; i < cap(channel); i++ {
		wg.Add(1)
		go printAndSleep(i, channel)
	}
	wg.Wait()
	close(channel)
	for number := range channel{
		values=append(values,number)
	}
	return
}

func printAndSleep(number int, channel chan int) {
	defer wg.Done()
	sleepTime := rand.Intn(3)
	time.Sleep(time.Duration(sleepTime)* time.Second)
	channel <- number
}
