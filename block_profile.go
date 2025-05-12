package main

import (
	"time"
)

func blockSimulationNoBlock() {
	ch := make(chan int, 1)
	ch <- 1
	<-ch
}

func blockSimulationWithBlock() {
	ch := make(chan int)
	go func() {
		time.Sleep(500 * time.Millisecond)
		ch <- 1
	}()
	<-ch
}
