package wyhash

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/zhangyunhao116/sbconv"
	"github.com/zhangyunhao116/wyhash/wyhashfv1"
)

func TestOptimizedFunc(t *testing.T) {
	for size := 0; size <= 257; size++ {
		data := make([]byte, size)
		for i := range data {
			data[i] = byte(rand.Intn(256))
		}
		if Sum64(data) != wyhashfv1.Sum64fv1(data) {
			t.Fatal(size, Sum64(data), wyhashfv1.Sum64fv1(data))
		}
		if Sum64String(sbconv.BytesToString(data)) != Sum64(data) {
			t.Fatal(size, Sum64(data), wyhashfv1.Sum64fv1(data))
		}
	}
}

func TestOptimizedWithSeedFunc(t *testing.T) {
	for size := 0; size <= 257; size++ {
		data := make([]byte, size)
		for i := range data {
			data[i] = byte(rand.Intn(256))
		}
		seed := uint64(rand.Int63())
		if Sum64WithSeed(data, seed) != wyhashfv1.Sum64fv1WithSeed(data, seed) {
			t.Fatal(size, Sum64(data), wyhashfv1.Sum64fv1(data))
		}
	}
}

func BenchmarkWyhash(b *testing.B) {
	sizes := []int{17, 21, 24, 29, 32,
		33, 64, 69, 96, 97, 128, 129, 240, 241,
		512, 1024, 100 * 1024,
	}

	for size := 0; size <= 16; size++ {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			b.SetBytes(int64(size))
			var (
				acc  uint64
				data = string(make([]byte, size))
			)
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				acc = Sum64(sbconv.StringToBytes(data))
			}
			runtime.KeepAlive(acc)
		})
	}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			b.SetBytes(int64(size))
			var acc uint64
			d := string(make([]byte, size))
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				acc = Sum64(sbconv.StringToBytes(d))
			}
			runtime.KeepAlive(acc)
		})
	}
}
