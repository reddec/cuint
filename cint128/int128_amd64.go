package cint128

import (
	"math/big"
)

//#include <stdint.h>
//
//typedef __int128 int128_t;
//typedef struct hl { uint64_t hi, lo; } hl;
//
//int128_t i128_add(int128_t x, int128_t y) {
//    return x + y;
//}
//
//int128_t i128_sub(int128_t x, int128_t y) {
//    return x - y;
//}
//
//
//int128_t i128_mul(int128_t x, int128_t y) {
//    return x * y;
//}
//
//int128_t i128_div(int128_t x, int128_t y) {
//    return x / y;
//}
//
//int128_t i128_from_i64(int64_t x) {
//    return (int128_t) (x);
//}
//
//int i128_sign(int128_t x) { return (x > 0) - (x < 0);}
//
//int i128_cmp(int128_t x, int128_t y) { return i128_sign(x - y);}
//
//hl i128_parts(int128_t x_) {
//    unsigned __int128 x = (unsigned __int128) x_;
//    uint64_t low  = (uint64_t)x;
//    uint64_t high = (x >> 64);
//    hl v; v.lo = low; v.hi = high;
//    return v;
//}
//
import "C"

type Int128 C.int128_t

func (u Int128) Add(x Int128) Int128 {
	return Int128(C.i128_add(C.int128_t(u), C.int128_t(x)))
}

func (u Int128) Sub(x Int128) Int128 {
	return Int128(C.i128_sub(C.int128_t(u), C.int128_t(x)))
}

func (u Int128) Mul(x Int128) Int128 {
	return Int128(C.i128_mul(C.int128_t(u), C.int128_t(x)))
}

func (u Int128) Div(x Int128) Int128 {
	return Int128(C.i128_div(C.int128_t(u), C.int128_t(x)))
}

func (u Int128) ToInt() *big.Int {
	var hl = C.i128_parts(C.int128_t(u))
	var v = new(big.Int).Lsh(new(big.Int).SetUint64(uint64(hl.hi)), 64)
	return new(big.Int).Add(v, new(big.Int).SetUint64(uint64(hl.lo)))
}

func (u Int128) Sign() int { return int(C.i128_sign(C.int128_t(u))) }

func (u Int128) Compare(x Int128) int {
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

func (u Int128) Equal(x Int128) bool {
	return ([16]byte(u)) == ([16]byte(x))
}

func (u Int128) Bytes() [16]byte { return C.int128_t(u) }

func (u *Int128) SetBytes(bts [16]byte) {
	*u = Int128(C.int128_t(bts))
}

func FromInt64(x int64) Int128 { return Int128(C.i128_from_i64(C.int64_t(x))) }
