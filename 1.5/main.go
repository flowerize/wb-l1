package main

import (
	"fmt"
	"time"
)

func main() {
	var N int
	fmt.Println("Type count of seconds: ")
	fmt.Scan(&N)
	timeout := time.After(time.Duration(N) * time.Second)

	dataCh := make(chan int)

	go func() {
		i := 0
		for {
			select {
			case <-timeout:
				close(dataCh)
				return
			default:
				dataCh <- i
				i++
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	go func() {
		for {
			select {
			case data := <-dataCh:
				fmt.Printf("Received: %d\n", data)
			case <-timeout:
				return
			}
		}
	}()

	<-timeout
	fmt.Println("Time's up. Exiting program.")
}
