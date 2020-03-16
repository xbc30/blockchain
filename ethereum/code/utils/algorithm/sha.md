## SHA

### SHA2 -> SHA256

#### GO语言使用SHA256
```go
//方法一
func digest1(data []byte)([]byte,error){
	h := sha256.Sum256(data)
	return h[:],nil
}

//方法二
func digest2(data []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(data)

	return h.Sum(nil),nil
}
```