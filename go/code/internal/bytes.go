package main

import (
	"bytes"
	"fmt"
)

func main(){
	// 转换
	b := []byte("seafood")  //强制类型转换
	a := bytes.ToUpper(b)

	fmt.Println(a, b)     //输出结果   [83 69 65 70 79 79 68] [115 101 97 102 111 111 100]

	// 比较
	c := []byte("asd")
	d := []byte("qwe")
	e := []byte("asd")

	fmt.Println(bytes.Compare(c, d)) // -1
	fmt.Println(bytes.Equal(c, e)) // true

	// 去除
	f := []byte("Hello World")

	fmt.Println(string(bytes.TrimRight(f, "World"))) // Hello

	// 分割 Split SplitAfter Fields
	splitSpace := []byte(" ")
	splitAfterBytes := bytes.Split(f, splitSpace)
	fmt.Println(string(splitAfterBytes[0])) // Hello
	fmt.Println(string(bytes.Join(splitAfterBytes, []byte("-")))) // Hello-World

	// 查找 HasPrefix Contains Index LastIndex
	fmt.Println(bytes.HasPrefix(f, []byte("H"))) // true

	// 替换
	g := []byte("Fuck")
	fmt.Println(string(bytes.Replace(f, []byte("Hello"), g, 1))) // Fuck World

	// 新建Buffer
	h := bytes.NewBuffer(g)
	fmt.Println(h.Cap()) // 4
	fmt.Println(string(h.Next(2))) // Fu
}
