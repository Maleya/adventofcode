package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func receivingChannel(inputChannel <-chan int, resultChannel chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	min := int(^uint(0) >> 1)

	for value := range inputChannel {
		if value < min {
			min = value
		}
	}

	// Send the minimum value to the resultChannel
	resultChannel <- min
}

func main() {
	rand.Seed(time.Now().UnixNano())

	inputChannel := make(chan int)
	resultChannel := make(chan int)

	// Create a WaitGroup for the receivingChannel
	var receivingWg sync.WaitGroup

	// Create a WaitGroup for the sending goroutine
	var sendWg sync.WaitGroup

	// Launch the receiving goroutine
	receivingWg.Add(1)
	go receivingChannel(inputChannel, resultChannel, &receivingWg)

	// Launch the sending goroutine
	sendWg.Add(1)
	go func() {
		defer sendWg.Done()
		defer close(inputChannel)

		for i := 0; i < 10; i++ {
			value := rand.Intn(100)
			fmt.Println("Received:", value)
			inputChannel <- value
			time.Sleep(time.Millisecond * 100)
		}
	}()

	// Use a separate goroutine to wait for both sendWg and receivingWg
	var waitAll sync.WaitGroup
	waitAll.Add(2)

	go func() {
		defer waitAll.Done()
		// Wait for the sending goroutine to finish
		sendWg.Wait()
		// Close the inputChannel after all values are sent
		close(inputChannel)
	}()

	go func() {
		defer waitAll.Done()
		// Wait for the receivingChannel goroutine to finish
		receivingWg.Wait()
		// Close the resultChannel after receiving is done
		close(resultChannel)
	}()

	// Wait for all goroutines to finish
	waitAll.Wait()

	// Retrieve the result from the result channel
	minimum, ok := <-resultChannel
	if ok {
		fmt.Println("Minimum value received:", minimum)
	} else {
		fmt.Println("Result channel closed.")
	}
}
