package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// 预定义变量, 保存命令行参数
	fmt.Println(os.Args)

	// 获取host name
	fmt.Println(os.Hostname())
	fmt.Println(os.Getpid())

	// 获取全部环境变量
	env := os.Environ()
	for k, v := range env {
		fmt.Println(k, v)
	}

	// 终止程序
	// os.Exit(1)

	// 获取一条环境变量
	fmt.Println(os.Getenv("PATH"))

	// 获取当前目录
	dir, err := os.Getwd()
	fmt.Println(dir, err)

	// 创建目录
	err = os.Mkdir(dir+"/new_file", 0755)
	fmt.Println(err)

	// 创建目录
	err = os.MkdirAll(dir+"/new", 0755)
	fmt.Println(err)

	// 删除目录
	err = os.Remove(dir + "/new_file")
	err = os.Remove(dir + "/new")
	fmt.Println(err)

	// 创建临时目录
	tmp_dir := os.TempDir()
	fmt.Println(tmp_dir)

	// 获取当前目录
	aDir, err := os.Getwd()
	fmt.Println(aDir, err)

	file := dir + "/new"
	var fh *os.File

	fi, _ := os.Stat(file)
	if fi == nil {
		fh, _ = os.Create(file) // 文件不存在就创建
	} else {
		fh, _ = os.OpenFile(file, os.O_RDWR, 0666) // 文件存在就打开
	}

	w := []byte("hello go language" + time.Now().String())
	n, err := fh.Write(w)
	fmt.Println(n, err)

	// 设置下次读写位置
	ret, err := fh.Seek(0, 0)
	fmt.Println("当前文件指针位置", ret, err)

	b := make([]byte, 128)
	n, err = fh.Read(b)
	fmt.Println(n, err, string(b))

	fh.Close()
}