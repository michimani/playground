package concurrency

import (
	"fmt"
	"time"
)

func Channel() {
	dataStream := make(chan interface{})

	go func() {
		dataStream <- "Hello channel!"
	}()

	fmt.Print(<-dataStream)
}

func BufferedChannel() {
	intStream := make(chan int, 3)
	// intStream := make(chan int)

	go func() {
		defer close(intStream)
		for i := 1; i <= 10; i++ {
			intStream <- i
			fmt.Println("added to channel. ", i)
		}
	}()

	for i := range intStream {
		time.Sleep(1 * time.Second)
		fmt.Println("receive from channel. ", i)
	}
}
