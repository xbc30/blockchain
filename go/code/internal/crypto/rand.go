package main


import (
	"fmt"
	"time"
	"math/rand"
)

func init(){
	//以时间作为初始化种子
	rand.Seed(time.Now().UnixNano())
}
func main() {

	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Println(a)
	}
	for i := 0; i < 10; i++ {
		a := rand.Intn(100)
		fmt.Println(a)
	}
	for i := 0; i < 10; i++ {
		a := rand.Float32()
		fmt.Println(a)
	}

}