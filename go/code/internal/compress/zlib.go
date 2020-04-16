package main

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

//进行zlib压缩
func DoZlibCompress(src []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	w.Write(src)
	w.Close()
	return in.Bytes()
}

//进行zlib解压缩
func DoZlibUnCompress(compressSrc []byte) []byte {
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r, _ := zlib.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}

func main() {
	buff := []byte{120, 156, 202, 72, 205, 201, 201, 215, 81, 40, 207,
		47, 202, 73, 225, 2, 4, 0, 0, 255, 255, 33, 231, 4, 147}
	b := bytes.NewReader(buff)
	r, err := zlib.NewReader(b)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, r)
	r.Close()

	zip := DoZlibCompress([]byte("hello, world\n"))
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("hello, world\n"))) // aGVsbG8sIHdvcmxkCg==
	fmt.Println(base64.StdEncoding.EncodeToString(zip)) // eJzKSM3JyddRKM8vyknhAgQAAP//IecEkw==
	fmt.Println(string(DoZlibUnCompress(zip))) // hello, world

	bigZip := DoZlibCompress([]byte("asdas6d415as1d2a12da12d1a21d5a1d5a15da51d5a151d5a15da51d5a15da1da5d15a\n"))
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("asdas6d415as1d2a12da12d1a21d5a1d5a15da51d5a151d5a15da51d5a15da1da5d15a\n"))) // YXNkYXM2ZDQxNWFzMWQyYTEyZGExMmQxYTIxZDVhMWQ1YTE1ZGE1MWQ1YTE1MWQ1YTE1ZGE1MWQ1YTE1ZGExZGE1ZDE1YQo===
	fmt.Println(base64.StdEncoding.EncodeToString(bigZip)) // eJxcx7EJgAAQxdDebf7hd59ANrj9QdTO4kHCyl6eKRuHjI8wsbwq/eK3Eqkpxx0AAP//+bQUeA==
	fmt.Println(string(DoZlibUnCompress(bigZip))) // asdas6d415as1d2a12da12d1a21d5a1d5a15da51d5a151d5a15da51d5a15da1da5d15a
}