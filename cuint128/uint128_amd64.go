package cuint128

import (
	"encoding/binary"
	"math/big"
)

//#include <stdint.h>
//
//typedef signed __int128 int128_t;
//typedef unsigned __int128 uint128_t;
//typedef struct hl { uint64_t hi, lo; } hl;
//
//uint128_t add(uint128_t x, uint128_t y) {
//    return x + y;
//}
//
//uint128_t sub(uint128_t x, uint128_t y) {
//    return x - y;
//}
//
//
//uint128_t mul(uint128_t x, uint128_t y) {
//    return x * y;
//}
//
//uint128_t div(uint128_t x, uint128_t y) {
//    return x / y;
//}
//
//uint128_t from_u64(uint64_t x) {
//    return (uint128_t) (x);
//}
//
//int sign(uint128_t x) { return (x > 0) - (x < 0);}
//
//int cmp(uint128_t x, uint128_t y) { return sign(x - y);}
//
//hl parts(uint128_t x) {
//    uint64_t low  = (uint64_t)x;
//    uint64_t high = (x >> 64);
//    hl v; v.lo = low; v.hi = high;
//    return v;
//}
//
import "C"

var hbo = binary.LittleEndian // native binary order for amd64

type UInt128 C.uint128_t

func (u UInt128) Add(x UInt128) UInt128 {
	return UInt128(C.add(C.uint128_t(u), C.uint128_t(x)))
}

func (u UInt128) Sub(x UInt128) UInt128 {
	return UInt128(C.sub(C.uint128_t(u), C.uint128_t(x)))
}

func (u UInt128) Mul(x UInt128) UInt128 {
	return UInt128(C.mul(C.uint128_t(u), C.uint128_t(x)))
}

func (u UInt128) Div(x UInt128) UInt128 {
	return UInt128(C.div(C.uint128_t(u), C.uint128_t(x)))
}

func (u UInt128) ToInt() *big.Int {
	var hl = C.parts(C.uint128_t(u))
	var v = new(big.Int).Lsh(new(big.Int).SetUint64(uint64(hl.hi)), 64)
	return new(big.Int).Add(v, new(big.Int).SetUint64(uint64(hl.lo)))
}

func (u UInt128) Sign() int { return int(C.sign(C.uint128_t(u))) }

func (u UInt128) Compare(x UInt128) int {
	var bts = [16]byte(u)
	var (
		uHi = hbo.Uint64(bts[:8])
		uLo = hbo.Uint64(bts[8:])
	)

	bts = [16]byte(x)
	var (
		oHi = hbo.Uint64(bts[:8])
		oLo = hbo.Uint64(bts[8:])
	)
	if uHi > oHi {
		return 1
	} else if uHi < oHi {
		return -1
	} else if uLo > oLo {
		return 1
	} else if uLo < oLo {
		return -1
	}
	return 0
}

func (u UInt128) Bytes() [16]byte { return C.uint128_t(u) }

func FromUInt64(x uint64) UInt128 { return UInt128(C.from_u64(C.uint64_t(x))) }
