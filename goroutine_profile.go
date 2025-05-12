package main

import (
	"time"
)

func goroutineSpawner() {
	for i := 0; i < 1000; i++ {
		go func() {
			time.Sleep(10 * time.Second)
		}()
	}
}
