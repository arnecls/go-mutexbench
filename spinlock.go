package main

import (
	"runtime"
	"sync/atomic"
)

type SpinLock struct {
	state *uint32
}

func NewSpinLock() SpinLock {
	return SpinLock{
		state: new(uint32),
	}
}

func (s *SpinLock) Lock() {
	for !atomic.CompareAndSwapUint32(s.state, 0, 1) {
		runtime.Gosched()
	}
}

func (s *SpinLock) Unlock() {
	atomic.StoreUint32(s.state, 0)
}
