
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//版权所有2017 Go Ethereum作者
//此文件是Go以太坊库的一部分。
//
//Go-Ethereum库是免费软件：您可以重新分发它和/或修改
//根据GNU发布的较低通用公共许可证的条款
//自由软件基金会，或者许可证的第3版，或者
//（由您选择）任何更高版本。
//
//Go以太坊图书馆的发行目的是希望它会有用，
//但没有任何保证；甚至没有
//适销性或特定用途的适用性。见
//GNU较低的通用公共许可证，了解更多详细信息。
//
//你应该收到一份GNU较低级别的公共许可证副本
//以及Go以太坊图书馆。如果没有，请参见<http://www.gnu.org/licenses/>。

package math

import (
	"bytes"
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestHexOrDecimal256(t *testing.T) {
	tests := []struct {
		input string
		num   *big.Int
		ok    bool
	}{
		{"", big.NewInt(0), true},
		{"0", big.NewInt(0), true},
		{"0x0", big.NewInt(0), true},
		{"12345678", big.NewInt(12345678), true},
		{"0x12345678", big.NewInt(0x12345678), true},
		{"0X12345678", big.NewInt(0x12345678), true},
//超前零行为测试：
{"0123456789", big.NewInt(123456789), true}, //注：不是八进制
		{"00", big.NewInt(0), true},
		{"0x00", big.NewInt(0), true},
		{"0x012345678abc", big.NewInt(0x12345678abc), true},
//无效语法：
		{"abcdef", nil, false},
		{"0xgg", nil, false},
//大于256位：
		{"115792089237316195423570985008687907853269984665640564039457584007913129639936", nil, false},
	}
	for _, test := range tests {
		var num HexOrDecimal256
		err := num.UnmarshalText([]byte(test.input))
		if (err == nil) != test.ok {
			t.Errorf("ParseBig(%q) -> (err == nil) == %t, want %t", test.input, err == nil, test.ok)
			continue
		}
		if test.num != nil && (*big.Int)(&num).Cmp(test.num) != 0 {
			t.Errorf("ParseBig(%q) -> %d, want %d", test.input, (*big.Int)(&num), test.num)
		}
	}
}

func TestMustParseBig256(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("MustParseBig should've panicked")
		}
	}()
	MustParseBig256("ggg")
}

func TestBigMax(t *testing.T) {
	a := big.NewInt(10)
	b := big.NewInt(5)

	max1 := BigMax(a, b)
	if max1 != a {
		t.Errorf("Expected %d got %d", a, max1)
	}

	max2 := BigMax(b, a)
	if max2 != a {
		t.Errorf("Expected %d got %d", a, max2)
	}
}

func TestBigMin(t *testing.T) {
	a := big.NewInt(10)
	b := big.NewInt(5)

	min1 := BigMin(a, b)
	if min1 != b {
		t.Errorf("Expected %d got %d", b, min1)
	}

	min2 := BigMin(b, a)
	if min2 != b {
		t.Errorf("Expected %d got %d", b, min2)
	}
}

func TestFirstBigSet(t *testing.T) {
	tests := []struct {
		num *big.Int
		ix  int
	}{
		{big.NewInt(0), 0},
		{big.NewInt(1), 0},
		{big.NewInt(2), 1},
		{big.NewInt(0x100), 8},
	}
	for _, test := range tests {
		if ix := FirstBitSet(test.num); ix != test.ix {
			t.Errorf("FirstBitSet(b%b) = %d, want %d", test.num, ix, test.ix)
		}
	}
}

func TestPaddedBigBytes(t *testing.T) {
	tests := []struct {
		num    *big.Int
		n      int
		result []byte
	}{
		{num: big.NewInt(0), n: 4, result: []byte{0, 0, 0, 0}},
		{num: big.NewInt(1), n: 4, result: []byte{0, 0, 0, 1}},
		{num: big.NewInt(512), n: 4, result: []byte{0, 0, 2, 0}},
		{num: BigPow(2, 32), n: 4, result: []byte{1, 0, 0, 0, 0}},
	}
	for _, test := range tests {
		if result := PaddedBigBytes(test.num, test.n); !bytes.Equal(result, test.result) {
			t.Errorf("PaddedBigBytes(%d, %d) = %v, want %v", test.num, test.n, result, test.result)
		}
	}
}

func BenchmarkPaddedBigBytesLargePadding(b *testing.B) {
	bigint := MustParseBig256("123456789123456789123456789123456789")
	for i := 0; i < b.N; i++ {
		PaddedBigBytes(bigint, 200)
	}
}

func BenchmarkPaddedBigBytesSmallPadding(b *testing.B) {
	bigint := MustParseBig256("0x18F8F8F1000111000110011100222004330052300000000000000000FEFCF3CC")
	for i := 0; i < b.N; i++ {
		PaddedBigBytes(bigint, 5)
	}
}

func BenchmarkPaddedBigBytesSmallOnePadding(b *testing.B) {
	bigint := MustParseBig256("0x18F8F8F1000111000110011100222004330052300000000000000000FEFCF3CC")
	for i := 0; i < b.N; i++ {
		PaddedBigBytes(bigint, 32)
	}
}

func BenchmarkByteAtBrandNew(b *testing.B) {
	bigint := MustParseBig256("0x18F8F8F1000111000110011100222004330052300000000000000000FEFCF3CC")
	for i := 0; i < b.N; i++ {
		bigEndianByteAt(bigint, 15)
	}
}

func BenchmarkByteAt(b *testing.B) {
	bigint := MustParseBig256("0x18F8F8F1000111000110011100222004330052300000000000000000FEFCF3CC")
	for i := 0; i < b.N; i++ {
		bigEndianByteAt(bigint, 15)
	}
}

func BenchmarkByteAtOld(b *testing.B) {

	bigint := MustParseBig256("0x18F8F8F1000111000110011100222004330052300000000000000000FEFCF3CC")
	for i := 0; i < b.N; i++ {
		PaddedBigBytes(bigint, 32)
	}
}

func TestReadBits(t *testing.T) {
	check := func(input string) {
		want, _ := hex.DecodeString(input)
		int, _ := new(big.Int).SetString(input, 16)
		buf := make([]byte, len(want))
		ReadBits(int, buf)
		if !bytes.Equal(buf, want) {
			t.Errorf("have: %x\nwant: %x", buf, want)
		}
	}
	check("000000000000000000000000000000000000000000000000000000FEFCF3F8F0")
	check("0000000000012345000000000000000000000000000000000000FEFCF3F8F0")
	check("18F8F8F1000111000110011100222004330052300000000000000000FEFCF3F8F0")
}

func TestU256(t *testing.T) {
	tests := []struct{ x, y *big.Int }{
		{x: big.NewInt(0), y: big.NewInt(0)},
		{x: big.NewInt(1), y: big.NewInt(1)},
		{x: BigPow(2, 255), y: BigPow(2, 255)},
		{x: BigPow(2, 256), y: big.NewInt(0)},
		{x: new(big.Int).Add(BigPow(2, 256), big.NewInt(1)), y: big.NewInt(1)},
//
		{x: big.NewInt(-1), y: new(big.Int).Sub(BigPow(2, 256), big.NewInt(1))},
		{x: big.NewInt(-2), y: new(big.Int).Sub(BigPow(2, 256), big.NewInt(2))},
		{x: BigPow(2, -255), y: big.NewInt(1)},
	}
	for _, test := range tests {
		if y := U256(new(big.Int).Set(test.x)); y.Cmp(test.y) != 0 {
			t.Errorf("U256(%x) = %x, want %x", test.x, y, test.y)
		}
	}
}

func TestBigEndianByteAt(t *testing.T) {
	tests := []struct {
		x   string
		y   int
		exp byte
	}{
		{"00", 0, 0x00},
		{"01", 1, 0x00},
		{"00", 1, 0x00},
		{"01", 0, 0x01},
		{"0000000000000000000000000000000000000000000000000000000000102030", 0, 0x30},
		{"0000000000000000000000000000000000000000000000000000000000102030", 1, 0x20},
		{"ABCDEF0908070605040302010000000000000000000000000000000000000000", 31, 0xAB},
		{"ABCDEF0908070605040302010000000000000000000000000000000000000000", 32, 0x00},
		{"ABCDEF0908070605040302010000000000000000000000000000000000000000", 500, 0x00},
	}
	for _, test := range tests {
		v := new(big.Int).SetBytes(common.Hex2Bytes(test.x))
		actual := bigEndianByteAt(v, test.y)
		if actual != test.exp {
			t.Fatalf("Expected  [%v] %v:th byte to be %v, was %v.", test.x, test.y, test.exp, actual)
		}

	}
}
func TestLittleEndianByteAt(t *testing.T) {
	tests := []struct {
		x   string
		y   int
		exp byte
	}{
		{"00", 0, 0x00},
		{"01", 1, 0x00},
		{"00", 1, 0x00},
		{"01", 0, 0x00},
		{"0000000000000000000000000000000000000000000000000000000000102030", 0, 0x00},
		{"0000000000000000000000000000000000000000000000000000000000102030", 1, 0x00},
		{"ABCDEF0908070605040302010000000000000000000000000000000000000000", 31, 0x00},
		{"ABCDEF0908070605040302010000000000000000000000000000000000000000", 32, 0x00},
		{"ABCDEF0908070605040302010000000000000000000000000000000000000000", 0, 0xAB},
		{"ABCDEF0908070605040302010000000000000000000000000000000000000000", 1, 0xCD},
		{"00CDEF090807060504030201ffffffffffffffffffffffffffffffffffffffff", 0, 0x00},
		{"00CDEF090807060504030201ffffffffffffffffffffffffffffffffffffffff", 1, 0xCD},
		{"0000000000000000000000000000000000000000000000000000000000102030", 31, 0x30},
		{"0000000000000000000000000000000000000000000000000000000000102030", 30, 0x20},
		{"ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 32, 0x0},
		{"ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 31, 0xFF},
		{"ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 0xFFFF, 0x0},
	}
	for _, test := range tests {
		v := new(big.Int).SetBytes(common.Hex2Bytes(test.x))
		actual := Byte(v, 32, test.y)
		if actual != test.exp {
			t.Fatalf("Expected  [%v] %v:th byte to be %v, was %v.", test.x, test.y, test.exp, actual)
		}

	}
}

func TestS256(t *testing.T) {
	tests := []struct{ x, y *big.Int }{
		{x: big.NewInt(0), y: big.NewInt(0)},
		{x: big.NewInt(1), y: big.NewInt(1)},
		{x: big.NewInt(2), y: big.NewInt(2)},
		{
			x: new(big.Int).Sub(BigPow(2, 255), big.NewInt(1)),
			y: new(big.Int).Sub(BigPow(2, 255), big.NewInt(1)),
		},
		{
			x: BigPow(2, 255),
			y: new(big.Int).Neg(BigPow(2, 255)),
		},
		{
			x: new(big.Int).Sub(BigPow(2, 256), big.NewInt(1)),
			y: big.NewInt(-1),
		},
		{
			x: new(big.Int).Sub(BigPow(2, 256), big.NewInt(2)),
			y: big.NewInt(-2),
		},
	}
	for _, test := range tests {
		if y := S256(test.x); y.Cmp(test.y) != 0 {
			t.Errorf("S256(%x) = %x, want %x", test.x, y, test.y)
		}
	}
}

func TestExp(t *testing.T) {
	tests := []struct{ base, exponent, result *big.Int }{
		{base: big.NewInt(0), exponent: big.NewInt(0), result: big.NewInt(1)},
		{base: big.NewInt(1), exponent: big.NewInt(0), result: big.NewInt(1)},
		{base: big.NewInt(1), exponent: big.NewInt(1), result: big.NewInt(1)},
		{base: big.NewInt(1), exponent: big.NewInt(2), result: big.NewInt(1)},
		{base: big.NewInt(3), exponent: big.NewInt(144), result: MustParseBig256("507528786056415600719754159741696356908742250191663887263627442114881")},
		{base: big.NewInt(2), exponent: big.NewInt(255), result: MustParseBig256("57896044618658097711785492504343953926634992332820282019728792003956564819968")},
	}
	for _, test := range tests {
		if result := Exp(test.base, test.exponent); result.Cmp(test.result) != 0 {
			t.Errorf("Exp(%d, %d) = %d, want %d", test.base, test.exponent, result, test.result)
		}
	}
}
