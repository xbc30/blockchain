package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main()  {
	// 1.字符串摘要
	// MD5
	h := md5.New()
	io.WriteString(h, "Hello World")
	fmt.Printf("%x\n", h.Sum(nil)) // b10a8db164e0754105b7a99be72e3fe5
	g := md5.New()
	g.Write([]byte("Hello World")) // 第二种输入方式
	fmt.Printf("%x\n", g.Sum(nil)) // 2b565b4301758d5a379dcf225b3d34f3
	f := g.Sum(nil) // 将字符串编码为16进制格式,返回字符串
	e := hex.EncodeToString(f)
	fmt.Println(e) // 2b565b4301758d5a379dcf225b3d34f3

	// SHA1
	a := sha1.New()
	io.WriteString(a, "Hello World")
	fmt.Printf("%x\n", a.Sum(nil)) // 0a4d55a8d778e5022fab701977c5d840bbc486d0

	// SHA256
	b := sha256.New()
	io.WriteString(b, "Hello World")
	fmt.Printf("%x\n", b.Sum(nil)) // a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e
	fmt.Printf("%x\n", b.Sum([]byte("Hello World"))) // a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e

	s := sha256.Sum256([]byte("Hello World"))
	fmt.Printf("%x\n", s) // a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e


	// SHA512
	c := sha512.New()
	io.WriteString(c, "Hello World")
	fmt.Printf("%x\n", c.Sum(nil)) // 2c74fd17edafd80e8447b0d46741ee243b7eb74dd2149a0ab1b9246fb30382f27e853d8585719e0e67cbda0daa8f51671064615d645ae27acb15bfb1447f459b

	// 2.文件摘要
	filePath := "./hash.txt"
	if hash , err := fileHash(filePath); err != nil {
		fmt.Printf(" %s, sha256 value: %s ", filePath,  hash) // ./hash.txt, sha256 hash: a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e

	}else {
		fmt.Printf(" %s, sha256 hash: %s ", filePath,  hash)
	}
}

func fileHash(filePath string) (string, error) {
	var hashValue string
	file, err := os.Open(filePath)
	if err != nil {
		return hashValue, err
	}
	defer file.Close()
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return hashValue,  err
	}
	hashInBytes := hash.Sum(nil)
	hashValue = hex.EncodeToString(hashInBytes)
	return hashValue, nil
}
