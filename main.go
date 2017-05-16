package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func Run(n int, f func()) {
	for i := 0; i < n; i++ {
		f()
	}
}

func Benchmark(name string, loops int, f func()) {
	t := time.Now()
	Run(loops, f)
	d := time.Since(t)
	fmt.Printf("%-15s : %f\n", name, float64(d.Nanoseconds())/1000000.0)
}

func BenchmarkThreaded(name string, loops int, f func()) {
	parallelN := runtime.NumCPU()
	started := new(sync.WaitGroup)
	ready := new(sync.WaitGroup)
	stopped := new(sync.WaitGroup)
	ready.Add(1)
	started.Add(parallelN)
	stopped.Add(parallelN)

	for i := 0; i < parallelN; i++ {
		go func() {
			started.Done()
			ready.Wait()
			Run(loops, f)
			stopped.Done()
		}()
	}

	started.Wait()
	ready.Done()

	t := time.Now()
	stopped.Wait()
	d := time.Since(t)

	fmt.Printf("%-15s : %f\n", name, float64(d.Nanoseconds())/1000000.0)
}

func MutexTest(loops int) {
	fmt.Println("Single threaded:")
	Benchmark("plain", loops, Calculation)
	Benchmark("mutex", loops, LockedWithMutex)
	Benchmark("mutex defer", loops, LockedWithDeferMutex)
	Benchmark("spinlock", loops, LockedWithSpinLock)
	Benchmark("spinlock defer", loops, LockedWithDeferSpinLock)

	fmt.Println("\nMulti threaded:")
	BenchmarkThreaded("plain", loops, Calculation)
	BenchmarkThreaded("mutex", loops, LockedWithMutex)
	BenchmarkThreaded("mutex defer", loops, LockedWithDeferMutex)
	BenchmarkThreaded("spinlock", loops, LockedWithSpinLock)
	BenchmarkThreaded("spinlock defer", loops, LockedWithDeferSpinLock)
}

func main() {
	MutexTest(1000000)
}
