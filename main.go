package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	//"runtime/trace"
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
	fmt.Printf("%-15s : %.2f\n", name, float64(d.Nanoseconds())/1000000.0)
}

func BenchmarkThreaded(name string, loops int, f func()) {
	parallelN := runtime.NumCPU() * 2
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

	fmt.Printf("%-15s : %.2f\n", name, float64(d.Nanoseconds())/1000000.0)
}

func MutexTest(loops int) {
	fmt.Println("Single threaded:")
	Benchmark("plain", loops, Calculation)
	Benchmark("mutex", loops, LockedWithMutex)
	Benchmark("mutex defer", loops, LockedWithDeferMutex)
	Benchmark("spinlock", loops, LockedWithSpinLock)
	Benchmark("spinlock defer", loops, LockedWithDeferSpinLock)
	Benchmark("spinlock2", loops, LockedWithSpinLock2)
	Benchmark("spinlock2 defer", loops, LockedWithDeferSpinLock2)

	fmt.Println("\nMulti threaded:")
	BenchmarkThreaded("plain", loops, Calculation)
	BenchmarkThreaded("mutex", loops, LockedWithMutex)
	BenchmarkThreaded("mutex defer", loops, LockedWithDeferMutex)
	BenchmarkThreaded("spinlock", loops, LockedWithSpinLock)
	BenchmarkThreaded("spinlock defer", loops, LockedWithDeferSpinLock)
	BenchmarkThreaded("spinlock2", loops, LockedWithSpinLock2)
	BenchmarkThreaded("spinlock defer2", loops, LockedWithDeferSpinLock2)
}

func main() {
	// Profile
	file, err := os.Create("cpu.profile")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := pprof.StartCPUProfile(file); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	// Trace
	/*traceFile, err := os.Create("run.trace")
	if err != nil {
		panic(err)
	}
	defer traceFile.Close()

	if err := trace.Start(traceFile); err != nil {
		panic(err)
	}
	defer trace.Stop()*/

	// Run test
	MutexTest(1000000)
}
