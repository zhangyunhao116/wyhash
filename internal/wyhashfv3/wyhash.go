// Package wyhashfv3 implements wyhash final version 3;
// DO NOT USE IT, for test only.
package wyhashfv3

import (
	"math/bits"
	"reflect"
	"unsafe"
)

const (
	s0 uint64 = 0xa0761d6478bd642f
	s1 uint64 = 0xe7037ed1a0b428db
	s2 uint64 = 0x8ebc6af09c88c6e3
	s3 uint64 = 0x589965cc75374cc3
	s4 uint64 = 0x1d8e4e27c47d124f
)

//go:linkname readUnaligned32 runtime.readUnaligned32
func readUnaligned32(p unsafe.Pointer) uint32

//go:linkname readUnaligned64 runtime.readUnaligned64
func readUnaligned64(p unsafe.Pointer) uint64

func _wyr8(x []byte) uint64 {
	array := *(*unsafe.Pointer)(unsafe.Pointer(&x))
	return readUnaligned64(array)
}

func _wyr4(x []byte) uint64 {
	array := *(*unsafe.Pointer)(unsafe.Pointer(&x))
	return uint64(readUnaligned32(array))
}

func _wyr3(x []byte, k int) uint64 {
	return uint64(x[0])<<16 | uint64(x[k>>1])<<8 | uint64(x[k-1])
}

func _wymix(a, b uint64) uint64 {
	hi, lo := bits.Mul64(a, b)
	return hi ^ lo
}

func Sum64(p []byte) uint64 {
	return Sum64WithSeed(p, s0)
}

func Sum64WithSeed(p []byte, seed uint64) uint64 {
	length := uint64(len(p))
	var a, b uint64
	if length <= 16 {
		if length >= 4 {
			// a=(_wyr4(p)<<32)|_wyr4(p+((len>>3)<<2)); b=(_wyr4(p+len-4)<<32)|_wyr4(p+len-4-((len>>3)<<2));
			a = _wyr4(p)<<32 | _wyr4(p[(length>>3)<<2:])
			b = _wyr4(p[length-4:])<<32 | _wyr4(p[length-4-((length>>3)<<2):])
		} else if length > 0 {
			// a=_wyr3(p,len); b=0;
			a = _wyr3(p, int(length))
		}
	} else {
		var (
			see1 = seed
			see2 = seed
		)
		// Note: the i will be len(p).
		if len(p) > 48 {
			for len(p) > 48 {
				// seed=_wymix(_wyr8(p)^secret[1],_wyr8(p+8)^seed);
				// see1=_wymix(_wyr8(p+16)^secret[2],_wyr8(p+24)^see1);
				// see2=_wymix(_wyr8(p+32)^secret[3],_wyr8(p+40)^see2);
				// p+=48; i-=48;
				seed = _wymix(_wyr8(p)^s1, _wyr8(p[8:])^seed)
				see1 = _wymix(_wyr8(p[16:])^s2, _wyr8(p[24:])^see1)
				see2 = _wymix(_wyr8(p[32:])^s3, _wyr8(p[40:])^see2)
				p = p[48:]
			}
			seed ^= see1 ^ see2
		}
		for len(p) > 16 {
			// seed=_wymix(_wyr8(p)^secret[1],_wyr8(p+8)^seed);  i-=16; p+=16;
			seed = _wymix(_wyr8(p)^s1, _wyr8(p[8:])^seed)
			p = p[16:]
		}
		// a=_wyr8(p+i-16);  b=_wyr8(p+i-8);
		// HACK
		var ap, bp []byte
		ph := (*reflect.SliceHeader)(unsafe.Pointer(&p))
		aph := (*reflect.SliceHeader)(unsafe.Pointer(&ap))
		bph := (*reflect.SliceHeader)(unsafe.Pointer(&bp))
		aph.Cap, aph.Len, bph.Cap, bph.Len = 16, 16, 16, 16
		aph.Data = uintptr(uint64(ph.Data) + uint64(len(p)) - 16)
		bph.Data = uintptr(uint64(ph.Data) + uint64(len(p)) - 8)
		a = _wyr8(ap)
		b = _wyr8(bp)
	}
	return _wymix(s1^length, _wymix(a^s1, b^seed))
}
