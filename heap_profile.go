package main

import (
	"strings"
	"time"
)

var globalLeak []*string // global slice holds all pointers, so GC can't collect them

// simulate memory leak function 1
func leakyFunction1() {
	for {
		// Profiling Result 1
		leak := make([]string, 0)
		leak = append(leak, "simulate memory leak")
		time.Sleep(100 * time.Millisecond)
	}
}

// memory leak function 2
func leakyFunction2() {
	for {
		/*
			Profiling Result 2
			- creates increasing memory heap that GC can't collect, to simulate a memory leak
		*/
		s := strings.Repeat("memory leak", 1000) // creates large strings to fill the heap
		ptr := &s                                // get pointer of each allocated string
		globalLeak = append(globalLeak, ptr)     // append pointer in global variable
		time.Sleep(10 * time.Millisecond)        // gives time for heap to grow
	}
}
