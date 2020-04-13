package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)
func main() {
	by := []byte{0x00, 0x00, 0x03, 0xe8}
	var num int32
	bytetoint(by, &num)
	fmt.Println(int(num), num)

	// 测试 int -> byte
	by2 := []byte{}
	var num2 int32
	num2 = 333
	by2 = inttobyte(&num2)
	fmt.Println(by2)

	// 测试 byte -> int
	var num3 int32
	bytetoint(by2,&num3)
	fmt.Println(num3)
}
// byte 转化 int
func bytetoint(by []byte, num *int32)  {
	b_buf := bytes.NewBuffer(by)
	binary.Read(b_buf, binary.BigEndian, num)
}
// 数字 转化 byte
func inttobyte(num *int32) []byte {
	b_buf := new(bytes.Buffer)
	binary.Write(b_buf, binary.BigEndian,num)
	return b_buf.Bytes()
}