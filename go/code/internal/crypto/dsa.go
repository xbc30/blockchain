package main

import (
	"crypto/dsa"
	"crypto/rand"
	"fmt"
)

func main() {
	// parameters 是私钥的参数
	var param dsa.Parameters
	// L1024N160是一个枚举，根据L1024N160来决定私钥的长度（L N）
	dsa.GenerateParameters(&param, rand.Reader, dsa.L1024N160)
	// 定义私钥的变量
	var privateKey dsa.PrivateKey
	// 设置私钥的参数
	privateKey.Parameters = param
	// 生成密钥对
	dsa.GenerateKey(&privateKey, rand.Reader)
	// 公钥是存在在私钥中的，从私钥中读取公钥
	publicKey := privateKey.PublicKey
	message := []byte("hello,dsa签名")

	// 进入签名操作
	r, s, _ := dsa.Sign(rand.Reader, &privateKey, message)

	fmt.Println(privateKey.X) // 303631090272549103387486975520499446626972635268

	fmt.Println(r) // 234949507860397206184759879738885880528592176291

	fmt.Println(s) // 501072401738848124963900257370706947203008030998

	// 进入验证
	flag := dsa.Verify(&publicKey, message, r, s)
	if flag {
		fmt.Println("数据未被修改")
	} else {
		fmt.Println("数据被修改")
	}
	flag = dsa.Verify(&publicKey, []byte("hello"), r, s)
	if flag {
		fmt.Println("数据未被修改")
	} else {
		fmt.Println("数据被修改")
	}
}