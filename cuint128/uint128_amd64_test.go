package cuint128

import (
	"fmt"
	"github.com/alecthomas/assert"
	"math/big"
	"testing"
)

func TestUInt128_Compare(t *testing.T) {

	var x = FromUInt64(1)
	var y = FromUInt64(2)

	fmt.Println(x.Compare(y))
	var i uint64
	for i = 0; i < uint64(100); i++ {
		var x = FromUInt64(1234567890 * i)
		var y = FromUInt64(5234567890 * i)

		var bx = new(big.Int).SetUint64(1234567890 * i)
		var by = new(big.Int).SetUint64(5234567890 * i)
		assert.True(t, x.Compare(y) == bx.Cmp(by))
	}

}
