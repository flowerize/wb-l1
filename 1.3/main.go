package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	var workersNum = flag.Int("n", 2, "Number of workers")
	var messagesNum = flag.Int("m", 10, "Number of messages")
	flag.Parse()

	if *workersNum <= 0 {
		fmt.Fprintf(os.Stderr, "Error: Number of workers must be > 0\n")
		flag.Usage()
		os.Exit(1)
	}
	if *messagesNum <= 0 {
		fmt.Fprintf(os.Stderr, "Error: Number of messages must be > 0\n")
		flag.Usage()
		os.Exit(1)
	}

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
