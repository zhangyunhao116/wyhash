package wyhash

import (
	"unsafe"

	"github.com/zhangyunhao116/sbconv"
	"github.com/zhangyunhao116/wyhash/internal/unalign"
)

func Sum64V3(data []byte) uint64 {
	return Sum64WithSeedV3(data, DefaultSeed)
}

func Sum64StringV3(data string) uint64 {
	return Sum64StringWithSeedV3(data, DefaultSeed)
}

func Sum64WithSeedV3(data []byte, seed uint64) uint64 {
	return Sum64StringWithSeedV3(sbconv.BytesToString(data), seed)
}

func Sum64StringWithSeedV3(data string, seed uint64) uint64 {
	var (
		a, b uint64
	)

	length := len(data)
	i := uintptr(len(data))
	paddr := *(*unsafe.Pointer)(unsafe.Pointer(&data))

	if length <= 16 {
		if length >= 4 {
			// Note: for short string, i == uintptr(length)
			// a=(_wyr4(p)<<32)|_wyr4(p+((len>>3)<<2)); b=(_wyr4(p+len-4)<<32)|_wyr4(p+len-4-((len>>3)<<2));
			a = unalign.Read4(paddr)<<32 | unalign.Read4(add(paddr, (i>>3)<<2))
			b = unalign.Read4(add(paddr, i-4))<<32 | unalign.Read4(add(paddr, i-4-((i>>3)<<2)))
		} else if length > 0 {
			// a=_wyr3(p,len); b=0;
			a = uint64(*(*byte)(paddr))<<16 | uint64(*(*byte)(add(paddr, uintptr(i>>1))))<<8 | uint64(*(*byte)(add(paddr, uintptr(i-1))))
		}
	} else {
		var (
			see1 = seed
			see2 = seed
		)
		if i > 48 {
			for i > 48 {
				// seed=_wymix(_wyr8(p)^secret[1],_wyr8(p+8)^seed);
				// see1=_wymix(_wyr8(p+16)^secret[2],_wyr8(p+24)^see1);
				// see2=_wymix(_wyr8(p+32)^secret[3],_wyr8(p+40)^see2);
				// p+=48; i-=48;
				seed = _wymix(unalign.Read8(paddr)^s1, unalign.Read8(add(paddr, 8))^seed)
				see1 = _wymix(unalign.Read8(add(paddr, 16))^s2, unalign.Read8(add(paddr, 24))^see1)
				see2 = _wymix(unalign.Read8(add(paddr, 32))^s3, unalign.Read8(add(paddr, 40))^see2)
				paddr = add(paddr, 48)
				i -= 48
			}
			seed ^= see1 ^ see2
		}
		for i > 16 {
			// seed=_wymix(unalign.Read8(p)^secret[1],unalign.Read8(p+8)^seed);  i-=16; p+=16;
			seed = _wymix(unalign.Read8(paddr)^s1, unalign.Read8(add(paddr, 8))^seed)
			paddr = add(paddr, 16)
			i -= 16
		}
		// a=unalign.Read8(p+i-16);  b=unalign.Read8(p+i-8);
		a = unalign.Read8(add(paddr, i-16))
		b = unalign.Read8(add(paddr, i-8))
	}
	// _wymix(secret[1]^len,_wymix(a^secret[1],b^seed));
	return _wymix(s1^uint64(length), _wymix(a^s1, b^seed))
}
