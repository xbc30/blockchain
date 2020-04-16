package main

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"log"
	"fmt"
	"io"
	"os"
)

func main() {
	// 一个缓冲区存储压缩的内容
	buf := bytes.NewBuffer(nil)

	// 创建一个flate.Write
	flateWrite, err := flate.NewWriterDict(buf, flate.BestCompression, []byte("key"))
	if err != nil {
		log.Fatalln(err)
	}
	defer flateWrite.Close()
	// 写入待压缩内容
	flateWrite.Write([]byte("compress/flate\n"))
	flateWrite.Flush()
	fmt.Println(base64.StdEncoding.EncodeToString(buf.Bytes())) // Ss7PLShKLS7WT8tJLEnlAgAAAP//

	// 解压刚压缩的内容
	flateReader := flate.NewReaderDict(buf, []byte("key"))
	defer flateReader.Close()
	// 输出
	io.Copy(os.Stdout, flateReader) // compress/flate
}