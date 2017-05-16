package main

import (
	"github.com/trivago/tgo/ttesting"
	"testing"
)

func TestSpinLockAsm(t *testing.T) {
	expect := ttesting.NewExpect(t)
	testValue := new(int32)

	result := Lock32(testValue)
	expect.True(result)
	expect.Equal(int32(1), *testValue)

	result = Lock32(testValue)
	expect.False(result)
	expect.Equal(int32(1), *testValue)

	Unlock32(testValue)
	expect.Equal(int32(0), *testValue)
}
