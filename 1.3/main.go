package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {
	var workersNum = flag.Int("n", 2, "Number of workers")
	var messagesNum = flag.Int("m", 10, "Number of messages")
	flag.Parse()

	ch := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < *workersNum; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for data := range ch {
				time.Sleep(200 * time.Millisecond)
				fmt.Printf("[Worker %d] Finished processing: %s\n", workerID, data)
			}
		}(i)
	}

	for i := 0; i < *messagesNum; i++ {
		ch <- fmt.Sprintf("Message #%d", i)
	}

	close(ch)

	wg.Wait()
}
