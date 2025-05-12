package main

import "time"

// CPU-intensive function 1
func cpuIntensive1() {
	for {
		for i := 0; i < 1000000; i++ {
			_ = i * i
		}

		// Profiling Result 1
		time.Sleep(time.Millisecond * 100)
	}
}

// CPU-intensive function 2
func cpuIntensive2() {
	for {
		for i := 0; i < 1000000; i++ {
			_ = i * i
		}

		// Profiling Result 2
		time.Sleep(time.Millisecond * 10)
	}
}
