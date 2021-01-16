package wyhash

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"testing"
	"unsafe"

	"github.com/zhangyunhao116/wyhash/wyhashfv1"
)

func TestAll(t *testing.T) {
	for size := 0; size <= 257; size++ {
		data := make([]byte, size)
		for i := range data {
			data[i] = byte(rand.Intn(256))
		}
		if Sum64Default(data) != wyhashfv1.Sum64fv1(data) {
			t.Fatal(size, Sum64Default(data), wyhashfv1.Sum64fv1(data))
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
				acc = Sum64Default(s2b(data))
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
				acc = Sum64Default(s2b(d))
			}
			runtime.KeepAlive(acc)
		})
	}
}

func s2b(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len
	return b
}
