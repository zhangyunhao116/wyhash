// Package wyhashfv1 implements wyhash final vesion 1.
// DO NOT USE IT, for test only.
package wyhashfv1

import (
	"math/bits"
	"unsafe"
)

const (
	s0 = 0xa0761d6478bd642f
	s1 = 0xe7037ed1a0b428db
	s2 = 0x8ebc6af09c88c6e3
	s3 = 0x589965cc75374cc3
	s4 = 0x1d8e4e27c47d124f
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

func _wyfinish16(p []byte, length int, seed uint64) uint64 {
	var a, b uint64
	i := len(p)
	if i < 8 {
		if i > 4 {
			a = _wyr4(p)
			b = _wyr4(p[i-4:])
		} else if i == 4 {
			a = _wyr4(p)
			b = 0
		} else if i != 0 {
			a = _wyr3(p, i)
			b = 0
		} else { // i == 0
			a, b = 0, 0
		}
	} else if i == 8 {
		a = _wyr8(p)
		b = 0
	} else {
		a = _wyr8(p)
		b = _wyr8(p[i-8:])
	}
	return _wymix(s1^uint64(length), _wymix(a^s1, b^seed))
}

func _wyfinish(p []byte, length int, seed uint64) uint64 {
	i := len(p)
	if i <= 16 {
		return _wyfinish16(p, length, seed)
	}
	// _wyfinish(p+16,len,_wymix(_wyr8(p)^secret[1],_wyr8(p+8)^seed),secret,i-16);
	return _wyfinish(p[16:], length, _wymix(_wyr8(p)^s1, _wyr8(p[8:])^seed))
}

func Sum64(p []byte) uint64 {
	var seed uint64 = s0
	length := len(p)
	if len(p) > 64 {
		var see1 = seed
		for len(p) > 64 {
			//   seed=_wymix(_wyr8(p)^secret[1], _wyr8(p+8)^seed)^_wymix(_wyr8(p+16)^secret[2], _wyr8(p+24)^seed);
			//   see1=_wymix(_wyr8(p+32)^secret[3],_wyr8(p+40)^see1)^_wymix(_wyr8(p+48)^secret[4],_wyr8(p+56)^see1);
			//   p+=64; i-=64;
			seed = _wymix(_wyr8(p)^s1, _wyr8(p[8:])^seed) ^ _wymix(_wyr8(p[16:])^s2, _wyr8(p[24:])^seed)
			see1 = _wymix(_wyr8(p[32:])^s3, _wyr8(p[40:])^see1) ^ _wymix(_wyr8(p[48:])^s4, _wyr8(p[56:])^see1)
			p = p[64:]
		}
		seed ^= see1
	}
	return _wyfinish(p, length, seed)
}

func Sum64WithSeed(p []byte, seed uint64) uint64 {
	length := len(p)
	if len(p) > 64 {
		var see1 = seed
		for len(p) > 64 {
			//   seed=_wymix(_wyr8(p)^secret[1], _wyr8(p+8)^seed)^_wymix(_wyr8(p+16)^secret[2], _wyr8(p+24)^seed);
			//   see1=_wymix(_wyr8(p+32)^secret[3],_wyr8(p+40)^see1)^_wymix(_wyr8(p+48)^secret[4],_wyr8(p+56)^see1);
			//   p+=64; i-=64;
			seed = _wymix(_wyr8(p)^s1, _wyr8(p[8:])^seed) ^ _wymix(_wyr8(p[16:])^s2, _wyr8(p[24:])^seed)
			see1 = _wymix(_wyr8(p[32:])^s3, _wyr8(p[40:])^see1) ^ _wymix(_wyr8(p[48:])^s4, _wyr8(p[56:])^see1)
			p = p[64:]
		}
		seed ^= see1
	}
	return _wyfinish(p, length, seed)
}
