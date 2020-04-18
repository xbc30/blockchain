package io

// https://segmentfault.com/a/1190000015591319

// 流式IO 文件IO 缓冲IO 网络IO 格式化IO(fmt.Printf)

import (
	"fmt"
	"io"
	"strings"
)

// alphaReader is a simple implementation of an io.Reader
// that streams only alpha chars from its string source.
// This example uses another reader as data source.
type alphaReader struct {
	reader io.Reader
}

func newAlphaReader(reader io.Reader) *alphaReader {
	return &alphaReader{reader: reader}
}

func alpha(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}

func (a *alphaReader) Read(p []byte) (int, error) {
	n, err := a.reader.Read(p)
	if err != nil {
		return n, err
	}
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		if char := alpha(p[i]); char != 0 {
			buf[i] = char
		}
	}

	copy(p, buf)
	return n, nil
}

func main() {
	// use an io.Reader as source for alphaReader
	reader := newAlphaReader(strings.NewReader("Hello! It's 9am, where is the sun?"))
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Print(string(p[:n]))
	}
	fmt.Println()
}