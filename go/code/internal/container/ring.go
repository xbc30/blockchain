package main

import (
	"container/ring"
	"fmt"
)

func main() {
	RingFunc()

}
func RingFunc() {
	r := ring.New(10) //初始长度10
	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}
	for i := 0; i < r.Len(); i++ {
		fmt.Println(r.Value)
		r = r.Next()
	}
	r = r.Move(6)
	fmt.Println(r.Value) //6
	r1 := r.Unlink(19)   //移除19%10=9个元素
	for i := 0; i < r1.Len(); i++ {
		fmt.Println(r1.Value)
		r1 = r1.Next()
	}
	fmt.Println(r.Len())  //10-9=1
	fmt.Println(r1.Len()) //9
}