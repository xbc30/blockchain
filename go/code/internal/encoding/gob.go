package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// encoding/gob包实现了高效的序列化，特别是数据结构较复杂的，结构体、数组和切片都被支持。

//定义一个结构体
type Student struct {
	Name string
	Age uint8
	Address string
}

func main(){
	//序列化
	s1:=Student{"张三",18,"江苏省"}
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)//创建编码器
	err1 := encoder.Encode(&s1)//编码
	if err1!=nil{
		fmt.Println(err1)
	}
	fmt.Printf("序列化后：%x\n",buffer.Bytes())

	//反序列化
	byteEn:=buffer.Bytes()
	decoder := gob.NewDecoder(bytes.NewReader(byteEn)) //创建解密器
	var s2 Student
	err2 := decoder.Decode(&s2)//解密
	if err2!=nil{
		fmt.Println(err2)
	}
	fmt.Println("反序列化后：",s2)
}