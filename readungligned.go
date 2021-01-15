// +build 386 amd64 arm arm64 ppc64le mips64le mipsle riscv64 wasm
//
// from golang-go/src/os/endian_big.go

package wyhash

import (
	"unsafe"
)

func _wyr8(p unsafe.Pointer) uint64 {
	q := (*[8]byte)(p)
	return uint64(q[0]) | uint64(q[1])<<8 | uint64(q[2])<<16 | uint64(q[3])<<24 | uint64(q[4])<<32 | uint64(q[5])<<40 | uint64(q[6])<<48 | uint64(q[7])<<56
}

func _wyr4(p unsafe.Pointer) uint64 {
	q := (*[4]byte)(p)
	return uint64(uint32(q[0]) | uint32(q[1])<<8 | uint32(q[2])<<16 | uint32(q[3])<<24)
}

// //go:linkname _wyr4 runtime.readUnaligned32
// func _wyr4(p unsafe.Pointer) uint32

// //go:linkname _wyr8 runtime.readUnaligned64
// func _wyr8(p unsafe.Pointer) uint64
