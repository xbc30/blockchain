package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	o := &sync.Once{}

	go do(o)
	go do(o)

	time.Sleep(time.Second * 2)
}



func do(o *sync.Once) {
	fmt.Println("Start do")

	o.Do(func() {
		fmt.Println("Doing something...")
	})

	fmt.Println("Do end")
}