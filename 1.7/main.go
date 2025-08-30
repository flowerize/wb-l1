package main

import (
	"fmt"
	"sync"
)

// потокобезопасная реализация map с синхронизацией через sync.Mutex
type ConcurrentMap struct {
	mu sync.Mutex
	m  map[string]int
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		m: make(map[string]int),
	}
}

func (cm *ConcurrentMap) Put(key string, value int) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.m[key] = value
}

func (cm *ConcurrentMap) Get(key string) (int, bool) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	val, exists := cm.m[key]
	return val, exists
}

func main() {
	cm := NewConcurrentMap()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", id)
			cm.Put(key, id*10)
			val, exists := cm.Get(key)
			if exists {
				fmt.Printf("Worker %d: %s = %d\n", id, key, val)
			}
		}(i)
	}
	wg.Wait()
}
