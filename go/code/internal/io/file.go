package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func writeFile(path string){
	//创建或截断打开文件
	f,err := os.Create(path)
	if err != nil{
		return
	}

	defer f.Close()

	f.WriteString("hello world\n")
}

func readFile(path string){
	//打开文件
	f,err := os.Open(path)
	if err != nil{
		return
	}
	defer f.Close()

	var buf []byte = make([]byte ,1024)
	n,err2 := f.Read(buf)
	if err2 != nil && err2 != io.EOF{
		fmt.Println(err2)
	}
	fmt.Println(n)
	fmt.Println(string(buf))
}

func readLine(path string){
	//打开文件
	f,err := os.Open(path)
	if err != nil{
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)		//带缓存 可以提高读写效率
	for{
		line,err := r.ReadBytes('\n')
		if err != nil{
			if err == io.EOF{
				break
			}
		}
		fmt.Println(string(line))
	}
}

func main() {
	url := "./a.txt"
	writeFile(url)
	readFile(url)
	readLine(url)
}