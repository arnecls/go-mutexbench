package main

import (
	"runtime"
	"sync/atomic"
)

type SpinLock struct {
	state *int64
}

func NewSpinLock() SpinLock {
	return SpinLock{
		state: new(int64),
	}
}

func (s *SpinLock) Lock() {
	for !atomic.CompareAndSwapInt64(s.state, 0, 1) {
		runtime.Gosched()
	}
}

func (s *SpinLock) Unlock() {
	atomic.StoreInt64(s.state, 0)
}
