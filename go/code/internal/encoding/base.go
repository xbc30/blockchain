package main

import "fmt"
import "encoding/hex"
import "encoding/base32"
import "encoding/base64"

func main() {
	s := "hello world!"

	sb := []byte(s)

	hexString := hex.EncodeToString(sb)
	hexByte, err := hex.DecodeString(hexString)
	fmt.Println(hexString)
	// 68656c6c6f20776f726c6421

	fmt.Println(hexByte, err)
	// [104 101 108 108 111 32 119 111 114 108 100 33] <nil>

	base32StdString := base32.StdEncoding.EncodeToString(sb)
	base32HexString := base32.HexEncoding.EncodeToString(sb)
	base32StdByte, err1 := base32.StdEncoding.DecodeString(base32StdString)
	base32HexByte, err2 := base32.HexEncoding.DecodeString(base32HexString)
	fmt.Println(base32StdString)
	// NBSWY3DPEB3W64TMMQQQ====

	fmt.Println(base32HexString)
	// D1IMOR3F41RMUSJCCGGG====

	fmt.Println(base32StdByte, err1)
	// [104 101 108 108 111 32 119 111 114 108 100 33] <nil>

	fmt.Println(base32HexByte, err2)
	// [104 101 108 108 111 32 119 111 114 108 100 33] <nil>

	base64StdString := base64.StdEncoding.EncodeToString(sb)
	base64UrlString := base64.URLEncoding.EncodeToString(sb)
	base64StdByte, err1 := base64.StdEncoding.DecodeString(base64StdString)
	base64UrlByte, err2 := base64.URLEncoding.DecodeString(base64UrlString)
	fmt.Println(base64StdString)
	// aGVsbG8gd29ybGQh

	fmt.Println(base64UrlString)
	// aGVsbG8gd29ybGQh

	fmt.Println(base64StdByte, err1)
	// [104 101 108 108 111 32 119 111 114 108 100 33] <nil>

	fmt.Println(base64UrlByte, err2)
	// [104 101 108 108 111 32 119 111 114 108 100 33] <nil>
}