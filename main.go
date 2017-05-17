package main

import (
	"fmt"
	"os"
	//"runtime/pprof"
	//"runtime/trace"
	"strconv"
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

func BenchmarkThreaded(name string, numRoutines, loops int, f func()) {
	started := new(sync.WaitGroup)
	ready := new(sync.WaitGroup)
	stopped := new(sync.WaitGroup)
	ready.Add(1)
	started.Add(numRoutines)
	stopped.Add(numRoutines)

	for i := 0; i < numRoutines; i++ {
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

func MutexTest(numRoutines, loops int) {
	fmt.Printf("Benchmarking %d loops\n", loops)
	if numRoutines < 2 {
		fmt.Println("Single threaded:")
		Benchmark("plain", loops, Calculation)
		Benchmark("mutex", loops, LockedWithMutex)
		Benchmark("spinlock", loops, LockedWithSpinLock)
		Benchmark("spinlock2", loops, LockedWithSpinLockBit)
		Benchmark("defer mutex", loops, LockedWithDeferMutex)
		Benchmark("defer spinlock", loops, LockedWithDeferSpinLock)
		Benchmark("defer spinlock2", loops, LockedWithDeferSpinLockBit)
	} else {
		fmt.Printf("Multi threaded (%d):\n", numRoutines)
		BenchmarkThreaded("plain", numRoutines, loops, Calculation)
		BenchmarkThreaded("mutex", numRoutines, loops, LockedWithMutex)
		BenchmarkThreaded("spinlock", numRoutines, loops, LockedWithSpinLock)
		BenchmarkThreaded("spinlock2", numRoutines, loops, LockedWithSpinLockBit)
		BenchmarkThreaded("defer mutex", numRoutines, loops, LockedWithDeferMutex)
		BenchmarkThreaded("defer spinlock", numRoutines, loops, LockedWithDeferSpinLock)
		BenchmarkThreaded("defer spinlock2", numRoutines, loops, LockedWithDeferSpinLockBit)
	}
}

func main() {
	if len(os.Args) != 3 {
		panic("two args required: num threads, num loops")
	}

	numRoutines, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("arg 1 (num go routines) is not a number")
	}

	loops, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic("arg 2 (num loops) is not a number")
	}

	// Profile
	/*file, err := os.Create("cpu.profile")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := pprof.StartCPUProfile(file); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	// Trace
	traceFile, err := os.Create("run.trace")
	if err != nil {
		panic(err)
	}
	defer traceFile.Close()

	if err := trace.Start(traceFile); err != nil {
		panic(err)
	}
	defer trace.Stop()*/

	// Run test
	MutexTest(numRoutines, loops)
}
