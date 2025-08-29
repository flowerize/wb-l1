package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// отправляет сообщение в канал и ловит панику
func safeSend(ch chan string, msg string) (ok bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered panic when sending: %v\n", r)
			ok = false
		}
	}()
	ch <- msg
	return true
}

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

	// Горутина для обработки сигнала
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

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

	sendDone := make(chan struct{})
	go func() {
		defer close(ch)
		for i := 0; i < *messagesNum; i++ {
			if !safeSend(ch, fmt.Sprintf("Message #%d", i)) {
				break
			}
		}
		sendDone <- struct{}{}
	}()

	select {
	case <-sigChan:
		fmt.Println("\n[Main] Received signal. Closing channel...")
		close(ch)
	case <-sendDone:
	}

	wg.Wait()
	fmt.Println("[Main] All workers finished. Exiting.")
}
