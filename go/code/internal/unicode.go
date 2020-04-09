package main

import "fmt"
import "unicode"

const (
	MaxRune = '\U0010FFFF' // Unicode 码点最大值
	ReplacementChar = '\uFFFD' // 代表无效的 Unicode 码点
	MaxASCII = '\u007F' // ASCII 码点最大值
	MaxLatin1 = '\u00FF' // Latin-1 码点最大值
)

func main() {
	a:= "Hello 世界"
	b:= "Hello World"

	// 判断字符 r 是否在 rangtab 表范围内
	for _, r := range a {
		// 判断字符是否为汉字
		if unicode.Is(unicode.Scripts["Han"], r) {
			fmt.Printf("%c", r) // 世界
		}
	}

	// 判断字符 r 是否为大小写格式
	for _, r:= range b {
		if unicode.IsUpper(r) {
			fmt.Printf("%c", r) // H W
		}
		if unicode.IsLower(r) {
			fmt.Printf("%c", r) // ello orld
		}
	}

	// To 将字符 r 转换为指定的格式
	for _, r := range a {
		fmt.Printf("%c", unicode.To(unicode.UpperCase, r))
	} // HELLO 世界 .ToUpper(r)
	for _, r := range a {
		fmt.Printf("%c", unicode.To(unicode.LowerCase, r))
	} // hello 世界 .ToLower(r)
	for _, r := range a {
		fmt.Printf("%c", unicode.To(unicode.TitleCase, r))
	} // HELLO 世界 .ToLower(r)

	// // IsDigit 判断 r 是否为一个十进制的数字字符
	c := "Hello 123１２３！"
	for _, r := range c {
		fmt.Printf("%c = %v\n", r, unicode.IsDigit(r))
	} // 123１２３ = true

	// IsGraphic IsPrint IsMark IsControl 少用

	// IsOneOf 判断 r 是否在 set 表范围内
	d := "Hello 世界！"
	// set 表设置为“汉字、标点符号”
	set := []*unicode.RangeTable{unicode.Han, unicode.P}
	for _, r := range d {
		fmt.Printf("%c = %v\n", r, unicode.IsOneOf(set, r))
	} // 世界！ = true

	// IsLetter 判断 r 是否为一个字母字符 (类别 L)
	e := "Hello\n\t世界！"
	for _, r := range e {
		fmt.Printf("%c = %v\n", r, unicode.IsLetter(r))
	} // Hello世界 = true

	// IsNumber 判断 r 是否为一个数字字符 (类别 N)
	f := "Hello 123１２３！"
	for _, r := range f {
		fmt.Printf("%c = %v\n", r, unicode.IsNumber(r))
	} // 123１２３ = true

	// IsSpace IsNumber ..
}