package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value int64
}

func (c *Counter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

func (c *Counter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

func main() {
	var counter Counter
	var wg sync.WaitGroup

	numGoroutines := 100
	iterationsPerGoroutine := 1000

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterationsPerGoroutine; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()

	fmt.Printf("Финальное значение счётчика: %d\n", counter.Value())
}
