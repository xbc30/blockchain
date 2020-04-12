package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	c := getSha256Code("test@example.com")
	fmt.Println(c)

	c = getHmacCode("test@example.com")
	fmt.Println(c)
}

func getHmacCode(s string) string {
	h := hmac.New(sha256.New, []byte("key"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil)) // 63a10897a46bca9755ba5edbfa6e079b3fa31d92e8d192731cb69b4d8ac4b00d
}

func getSha256Code(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil)) // 973dfe463ec85785f5f95af5ba3906eedb2d931c24e69824a89ea65dba4e813b
}
