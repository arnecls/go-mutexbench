package main

import (
	"runtime"
)

type SpinLockBit struct {
	state *int32
}

func NewSpinLockBit() SpinLockBit {
	return SpinLockBit{
		state: new(int32),
	}
}

func (s *SpinLockBit) Lock() {
	for !TestAndSet32(s.state) {
		runtime.Gosched()
	}
}

func (s *SpinLockBit) Unlock() {
	TestAndReset32(s.state)
}
