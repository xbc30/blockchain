package main

import (
	"crypto/des"
	"bytes"
	"crypto/cipher"
	"fmt"
	"encoding/base64"
)

//DES加密方法
func MyDESEncrypt (origData,key []byte){
	//将字节秘钥转换成block快
	block,_ := des.NewCipher(key)
	//对明文先进行补码操作
	origData = PKCS5Padding(origData,block.BlockSize())
	//设置加密方式
	blockMode := cipher.NewCBCEncrypter(block,key)
	//创建明文长度的字节数组
	crypted := make([]byte, len(origData))
	//加密明文,加密后的数据放到数组中
	blockMode.CryptBlocks(crypted,origData)
	//将字节数组转换成字符串
	fmt.Println(base64.StdEncoding.EncodeToString(crypted))
}
//实现明文的补码
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	//计算出需要补多少位
	padding := blockSize - len(ciphertext)%blockSize
	//Repeat()函数的功能是把参数一 切片复制 参数二count个,然后合成一个新的字节切片返回
	// 需要补padding位的padding值
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	//把补充的内容拼接到明文后面
	return append(ciphertext,padtext...)
}

//解密
func MyDESDecrypt(data string, key []byte) {
	//倒叙执行一遍加密方法
	//将字符串转换成字节数组
	crypted,_ := base64.StdEncoding.DecodeString(data)
	//将字节秘钥转换成block快
	block, _ := des.NewCipher(key)
	//设置解密方式
	blockMode := cipher.NewCBCDecrypter(block,key)
	//创建密文大小的数组变量
	origData := make([]byte, len(crypted))
	//解密密文到数组origData中
	blockMode.CryptBlocks(origData,crypted)
	//去补码
	origData = PKCS5UnPadding(origData)
	//打印明文
	fmt.Println(string(origData))
}

//去除补码
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	//解密去补码时需取最后一个字节，值为m，则从数据尾部删除m个字节，剩余数据即为加密前的原文
	return origData[:(length - unpadding)]
}
func main(){
	//定义明文
	data := []byte("hello world")
	//密钥
	key := []byte("12345678")
	//加密
	MyDESEncrypt(data,key) // CyqS6B+0nOGkMmaqyup7gQ==
	//解密
	MyDESDecrypt("CyqS6B+0nOGkMmaqyup7gQ==",key) // hello world
}