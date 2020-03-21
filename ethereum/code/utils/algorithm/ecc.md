## ECC

**GO语言使用ECC:**

```go
import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"os"
	"crypto/x509"
	"encoding/pem"
	"crypto/sha256"
	"math/big"
	"fmt"
)
//生成ECC椭圆曲线密钥对，保存到文件
func GenerateECCKey() {
	//生成密钥对
	privateKey, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	if err != nil {
		panic(err)
	}
	//保存私钥
	//生成文件
	privatefile, err := os.Create("eccprivate.pem")
	if err != nil {
		panic(err)
	}
	//x509编码
	eccPrivateKey, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		panic(err)
	}
	//pem编码
	privateBlock := pem.Block{
		Type:  "ecc private key",
		Bytes: eccPrivateKey,
	}
	pem.Encode(privatefile, &privateBlock)
	//保存公钥
	publicKey := privateKey.PublicKey
	//创建文件
	publicfile, err := os.Create("eccpublic.pem")
	//x509编码
	eccPublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	//pem编码
	block := pem.Block{Type: "ecc public key", Bytes: eccPublicKey}
	pem.Encode(publicfile, &block)
}

//取得ECC私钥
func GetECCPrivateKey(path string) *ecdsa.PrivateKey {
	//读取私钥
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	//pem解码
	block, _ := pem.Decode(buf)
	//x509解码
	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	return privateKey
}

//取得ECC公钥
func GetECCPublicKey(path string) *ecdsa.PublicKey {
	//读取公钥
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	//pem解密
	block, _ := pem.Decode(buf)
	//x509解密
	publicInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	publicKey := publicInterface.(*ecdsa.PublicKey)
	return publicKey
}

//对消息的散列值生成数字签名
func SignECC(msg []byte, path string)([]byte,[]byte) {
	//取得私钥
	privateKey := GetECCPrivateKey(path)
	//计算哈希值
	hash := sha256.New()
	//填入数据
	hash.Write(msg)
	bytes := hash.Sum(nil)
	//对哈希值生成数字签名
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, bytes)
	if err != nil {
		panic(err)
	}
	rtext, _ := r.MarshalText()
	stext, _ := s.MarshalText()
	return rtext, stext
}

//验证数字签名
func VerifySignECC(msg []byte,rtext,stext []byte,path string) bool{
	//读取公钥
	publicKey:=GetECCPublicKey(path)
	//计算哈希值
	hash := sha256.New()
	hash.Write(msg)
	bytes := hash.Sum(nil)
	//验证数字签名
	var r,s big.Int
	r.UnmarshalText(rtext)
	s.UnmarshalText(stext)
	verify := ecdsa.Verify(publicKey, bytes, &r, &s)
	return verify
}
//测试
func main() {
	//生成ECC密钥对文件
	GenerateECCKey()

	//模拟发送者
	//要发送的消息
	msg:=[]byte("hello world")
	//生成数字签名
	rtext,stext:=SignECC(msg,"eccprivate.pem")

	//模拟接受者
	//接受到的消息
	acceptmsg:=[]byte("hello world")
	//接收到的签名
	acceptrtext:=rtext
	acceptstext:=stext
	//验证签名
	verifySignECC := VerifySignECC(acceptmsg, acceptrtext, acceptstext, "eccpublic.pem")
	fmt.Println("验证结果：",verifySignECC)
}
```

**ECDSA**
> 椭圆曲线数字签名算法（ECDSA）是使用椭圆曲线密码（ECC）对数字签名算法（DSA）的模拟

```go
package main

import (
    "crypto/ecdsa"
    "crypto/rand"
    "fmt"
    "crypto/elliptic"
    "log"
)

func main() {
    // 生成公钥和私钥
    privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    if err != nil {
        log.Fatalln(err)
    }
    // 公钥是存在在私钥中的，从私钥中读取公钥
    publicKey := &privateKey.PublicKey
    message := []byte("hello,dsa签名")

    // 进入签名操作
    r, s, _ := ecdsa.Sign(rand.Reader, privateKey, message)
    // 进入验证
    flag := ecdsa.Verify(publicKey, message, r, s)
    if flag {
        fmt.Println("数据未被修改")
    } else {
        fmt.Println("数据被修改")
    }
    flag = ecdsa.Verify(publicKey, []byte("hello"), r, s)
    if flag {
        fmt.Println("数据未被修改")
    } else {
        fmt.Println("数据被修改")
    }
}
```

**Secp256k1和Secp256r1:**
> Secp256k1是指比特币中使用的ECDSA(椭圆曲线数字签名算法)曲线的参数, ECC有多个参数用来调节速度和安全性，比特币和以太坊使用 secp256k1参数。Secp256k1和Secp256r1两者都是场 zp 上的椭圆曲线，其中 p 是256位素数（尽管每条曲线有不同的素数）。secp256k1中的 “k” 代表 Koblitz，sepc256r1中 的 “r” 代表 随机。Koblitz椭圆曲线具有一些特殊属性，可以更有效地实现组操作。据信存在小的安全权衡，更 “随机” 选择的参数更安全。然而，有些人怀疑随机系数可能已经被选择来提供后门。
                                                                                     
