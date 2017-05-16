package main

import (
	"runtime"
	"sync/atomic"
)

func Lock32(v *int32) bool
func Unlock32(v *int32)

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

type SpinLock2 struct {
	state *int32
}

func NewSpinLock2() SpinLock2 {
	return SpinLock2{
		state: new(int32),
	}
}

func (s *SpinLock2) Lock() {
	for !Lock32(s.state) {
		runtime.Gosched()
	}
}

func (s *SpinLock2) Unlock() {
	Unlock32(s.state)
}
