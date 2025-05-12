package main

import (
	"sync"
	"time"
)

var mu sync.Mutex

func simulateLowContention() {
	for i := 0; i < 100; i++ {
		mu.Lock()
		time.Sleep(1 * time.Millisecond)
		mu.Unlock()
	}
}

func simulateHighContention() {
	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			time.Sleep(10 * time.Millisecond)
			mu.Unlock()
		}()
	}
	time.Sleep(2 * time.Second)
}
