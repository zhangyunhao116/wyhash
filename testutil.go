package wyhash

import (
	"math/rand"
	"testing"
)

func testSum(t *testing.T, f1, f2 func(p []byte) uint64) {
	for size := 0; size <= 257; size++ {
		data := make([]byte, size)
		for i := range data {
			data[i] = byte(rand.Intn(256))
		}
		if f1(data) != f2(data) {
			t.Fatal(size, f1(data), f2(data))
		}
	}
}

func testSumWithSeed(t *testing.T, f1, f2 func(p []byte, seed uint64) uint64) {
	for size := 0; size <= 257; size++ {
		data := make([]byte, size)
		for i := range data {
			data[i] = byte(rand.Intn(256))
		}
		seed := uint64(rand.Int63())
		if f1(data, seed) != f2(data, seed) {
			t.Fatal(size, f1(data, seed), f2(data, seed))
		}
	}
}
