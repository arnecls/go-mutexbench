package main

import (
	"math/rand"
	"sync"
)

var globalState int64
var mutex *sync.Mutex
var spinLock SpinLock

func init() {
	mutex = new(sync.Mutex)
	spinLock = NewSpinLock()
}

func Calculation() {
	globalState += rand.Int63n(1000)
}

func LockedWithMutex() {
	mutex.Lock()
	Calculation()
	mutex.Unlock()
}

func LockedWithSpinLock() {
	spinLock.Lock()
	Calculation()
	spinLock.Unlock()
}

func LockedWithDeferMutex() {
	mutex.Lock()
	defer mutex.Unlock()
	Calculation()
}

func LockedWithDeferSpinLock() {
	spinLock.Lock()
	defer spinLock.Unlock()
	Calculation()
}
