## RLP
> RLP（Recursive Length Rrefix, 递归长度前缀）提供了一种适用于任意二进制数据数组的编码，RLP已经称为以太坊中对对象进行序列化的主要编码方式。RLP的唯一目标就是解决结构体的编码问题；对原子数据类型（比如：字符串，整数型，浮点型）的编码则交给更高层的协议；以太坊中要求数字必须是一个大端字节序的、没有零占位的存储的格式(也就是说，一个整数0和一个空数组是等同的).RLP目的是可以将常用的数据结构,uint,string,[]byte,struct,slice,array,big.int等序列化以及反序列化.与其他序列化方法相比，RLP编码的优点在于使用了灵活的长度前缀来表示数据的实际长度，并且使用递归的方式能编码相当大的数据，当接收或者解码经过RLP编码后的数据时，根据第1个字节就能推断数据的类型、大概长度和数据本身等信息。而其他的序列化方法， 不能根据第1个字节获得如此多的信息量                                                                                                                                                                                                                                                                                                                                                                                                                     
  
**数据定义:**  

RLP编码的定义只处理以下2类底层数据：

* 字符串（string）是指字节数组。例如，空串”“，再如单词”cat”，以及句子”Lorem ipsum dolor sit amet, consectetur adipisicing elit”等。

* 列表（list）是一个可嵌套结构，里面可包含字符串和列表。例如，空列表[]，再如一个包含两个字符串的列表[“cat”,”dog”]，再比如嵌套列表的复杂列表[“cat”, [“puppy”, “cow”], “horse”, [[]], “pig”, [“”], “sheep”]。

所有上层类型的数据需要转成以上的2类数据，才能进行RLP编码。转换的规则RLP编码不统一规定，可以自定义转换规则。例如struct可以转成列表；int可以转成二进制序列（属于字符串这一类, 必须去掉首部0，必须用大端模式表示）；map类型可以转换为由k和v组成的结构体、k按字典顺序排列的列表：[[k1,v1],[k2,v2]…] 等
  
**五个规则:**

规则 | 规则一 | 规则二 | 规则三 | 规则四 | 规则五  
---- | ---- | ---- | ---- | ---- | ----
范围 | 0x00~0x7f | 0x80~0xb7 | 0xb8~0xbf | 0xc0~0xf7 | 0xf8~0xff 

* 规则一
```javascript
考虑到ASCII码的特殊性，如果输入的是ASCII码（值范围[0x00, 0x7f])，他的RLP编码就是他本身
```

* 规则二
```javascript
如果输入的长度是0-55字节，他的RLP编码包含一个单字节的前缀，后面跟着字符串本身，这个前缀的值是0x80加上字符串的长度。
由于被编码的字符串最大长度是55=0x37,因此单字节前缀的最大值是0x80+0x37=0xb7，
即编码的第一个字节的取值范围是[0x80, 0xb7]。（长度范围为什么是0-55，看完规则三就知道了
```

* 规则三
```javascript
如果字符串的长度大于55个字节，它的RLP编码包含一个单字节的前缀，后面跟着字符串的长度，在后面跟着字符串本身。
这个前缀的值是0xb7加上字符串长度的二进制形式的字节长度，说的有点绕，举个例子就明白了，例如一个字符串的长度是1024，
它的二进制形式是10000000000，这个二进制形式的长度是2个字节，所以前缀应该是0xb7+2=0xb9，字符串长度1024=0x400，
因此整个RLP编码应该是\xb9\x04\x00再跟上字符串本身。编码的第一个字节即前缀的取值范围是[0xb8, 0xbf]，
因为字符串长度二进制形式最少是1个字节，因此最小值是0xb7+1=0xb8，字符串长度二进制最大是8个字节，因此最大值是0xb7+8=0xbf。
（支持的最大长度是8个字节，2^32大小也够用了。另外规则二中为什么最大长度是55，因为55+8=64 -1，范围最大值刚好是0xbf)
```

* 规则四
```javascript
如果一个列表的总长度（列表的总长度指的是它包含的项的个数加上它包含的各项的长度之和）是0-55字节，
它的RLP编码包含一个单字节的前缀，后面跟着列表中各元素的RLP编码，这个前缀的值是0xc0加上列表的总长度。
编码的第一个字节的取值范围是[0xc0, 0xf7].
```

* 规则五
```javascript
如果一个列表的总长度大于55字节，它的RLP编码包含一个单字节的前缀，后面跟着列表的长度，后面再跟着列表中各元素项的RLP编码，
这个前缀的值是0xf7加上列表总长度的二进制形式的字节长度。编码的第一个字节的取值范围是[0xf8, 0xff]

```

**关键函数:**
* encode
```go
func Encode(w io.Writer, val interface{}) error {
	if outer, ok := w.(*encbuf); ok {
		// Encode was called by some type's EncodeRLP.
		// Avoid copying by writing to the outer encbuf directly.
		return outer.encode(val)
	}
	eb := encbufPool.Get().(*encbuf)
	defer encbufPool.Put(eb)
	eb.reset()
	if err := eb.encode(val); err != nil {
		return err
	}
	return eb.toWriter(w)
}

// EncodeToBytes returns the RLP encoding of val.
// Please see the documentation of Encode for the encoding rules.
func EncodeToBytes(val interface{}) ([]byte, error) {
	eb := encbufPool.Get().(*encbuf)
	defer encbufPool.Put(eb)
	eb.reset()
	if err := eb.encode(val); err != nil {
		return nil, err
	}
	return eb.toBytes(), nil
}

// EncodeToReader returns a reader from which the RLP encoding of val
// can be read. The returned size is the total size of the encoded
// data.
//
// Please see the documentation of Encode for the encoding rules.
func EncodeToReader(val interface{}) (size int, r io.Reader, err error) {
	eb := encbufPool.Get().(*encbuf)
	eb.reset()
	if err := eb.encode(val); err != nil {
		return 0, nil, err
	}
	return eb.size(), &encReader{buf: eb}, nil
}
```

* decode

```go
func Decode(r io.Reader, val interface{}) error {
	stream := streamPool.Get().(*Stream)
	defer streamPool.Put(stream)

	stream.Reset(r, 0)
	return stream.Decode(val)
}

// DecodeBytes parses RLP data from b into val.
// Please see the documentation of Decode for the decoding rules.
// The input must contain exactly one value and no trailing data.
func DecodeBytes(b []byte, val interface{}) error {
	r := bytes.NewReader(b)
	stream := streamPool.Get().(*Stream)
	defer streamPool.Put(stream)

	stream.Reset(r, uint64(len(b)))
	if err := stream.Decode(val); err != nil {
		return err
	}
	if r.Len() > 0 {
		return ErrMoreThanOneValue
	}
	return nil
}
```

**例子:**
```go
type TestRlpStruct struct {
    A      uint
    B      string
    C      []byte
    BigInt *big.Int
}

//rlp用法
func TestRlp(t *testing.T) {
    //1.将一个整数数组序列化
    arrdata, err := rlp.EncodeToBytes([]uint{32, 28})
    fmt.Printf("unuse err:%v\n", err)
    //fmt.Sprintf("data=%s,err=%v", hex.EncodeToString(arrdata), err)
    //2.将数组反序列化
    var intarray []uint
    err = rlp.DecodeBytes(arrdata, &intarray)
    //intarray 应为{32,28}
    fmt.Printf("intarray=%v\n", intarray)

    //3.将一个布尔变量序列化到一个writer中
    writer := new(bytes.Buffer)
    err = rlp.Encode(writer, true)
    //fmt.Sprintf("data=%s,err=%v",hex.EncodeToString(writer.Bytes()),err)
    //4.将一个布尔变量反序列化
    var b bool
    err = rlp.DecodeBytes(writer.Bytes(), &b)
    //b:true
    fmt.Printf("b=%v\n", b)

    //5.将任意一个struct序列化
    //将一个struct序列化到reader中
    _, r, err := rlp.EncodeToReader(TestRlpStruct{3, "44", []byte{0x12, 0x32}, big.NewInt(32)})
    var teststruct TestRlpStruct
    err = rlp.Decode(r, &teststruct)
    //{A:0x3, B:"44", C:[]uint8{0x12, 0x32}, BigInt:32}
    fmt.Printf("teststruct=%#v\n", teststruct)

}
```
