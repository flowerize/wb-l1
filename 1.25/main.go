package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	start := time.Now()
	for {
		if time.Since(start) >= duration {
			break
		}
	}
}

func main() {
	fmt.Println("Start")
	sleep(3 * time.Second)
	fmt.Println("End")
}
