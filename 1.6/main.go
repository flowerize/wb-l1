package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Exit by condition ===")
	exitByCondition()

	fmt.Println("\n=== Exit by channel notification ===")
	exitByChannel()

	fmt.Println("\n=== Exit by context ===")
	exitByContext()

	fmt.Println("\n=== Exit by runtime.Goexit ===")
	exitByGoexit()

	fmt.Println("\n=== Exit by closing a channel ===")
	exitByClosingChannel()

	fmt.Println("\nAll examples completed.")
}

// 1. Выход по условию
func exitByCondition() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Printf("Condition check: %d\n", i)
			time.Sleep(200 * time.Millisecond)
		}
		fmt.Println("Exit by condition check.")
	}()

	wg.Wait()
}

// 2. Выход через канал уведомления
func exitByChannel() {
	stop := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-stop:
				fmt.Println("Exit by channel signal.")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	close(stop)
	wg.Wait()
}

// 3. Выход через контекст (context.Context)
func exitByContext() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Exit by context timeout.")
				return
			default:
				fmt.Println("Context-based work...")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	wg.Wait()
}

// 4. Выход через runtime.Goexit()
func exitByGoexit() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer fmt.Println("Deferred in Goexit example.")
		for i := 0; i < 5; i++ {
			if i == 3 {
				fmt.Println("Forced exit via runtime.Goexit().")
				runtime.Goexit()
			}
			fmt.Printf("Goexit loop: %d\n", i)
			time.Sleep(200 * time.Millisecond)
		}
	}()

	time.Sleep(1 * time.Second)
	wg.Wait()
}

// 5. Выход через закрытие канала
func exitByClosingChannel() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for data := range ch {
			fmt.Printf("Received: %d\n", data)
		}
		fmt.Println("Exit by closing channel.")
	}()

	for i := 0; i < 3; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()
}
