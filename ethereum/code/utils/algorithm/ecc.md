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