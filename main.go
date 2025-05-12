package main

import (
	"log"
	"net/http"
	_ "net/http/pprof" // Import pprof
	"runtime"
)

// Configuration for profiling
func enableProfiling() {
	// Enable mutex and block profiling
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	// Start pprof server
	go func() {
		log.Println("Starting pprof server on :6060")
		http.ListenAndServe("localhost:6060", nil)
	}()
}

func main() {
	// Enable pprof
	enableProfiling()

	// CPU Profile
	go cpuIntensive1() // uses time.Sleep(100ms)
	go cpuIntensive2() // time.Sleep(10ms) for more CPU usage

	// Heap Profile
	go leakyFunction1() // no actual leak
	go leakyFunction2() // simulated memory leak (heap grows over time)

	// Goroutine Profile
	goroutineSpawner()

	// Block Profile
	blockSimulationNoBlock()
	blockSimulationWithBlock()

	// Mutex Profile
	simulateLowContention()
	go simulateHighContention()

	// Keep main alive for profiling
	select {}
}
