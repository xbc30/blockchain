package main


import (
	"fmt"
	"sync"
	"time"
)


var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)


func test(x int) {
	cond.L.Lock() //获取锁
	fmt.Println("aaa: ", x)
	cond.Wait() //等待通知  暂时阻塞
	fmt.Println("bbb: ", x)
	time.Sleep(time.Second * 2)
	cond.L.Unlock()
}


func main() {
	for i := 0; i < 5; i++ {
		go test(i)
	}
	fmt.Println("start all")
	time.Sleep(time.Second * 1)
	fmt.Println("broadcast")
	cond.Signal()// 下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 1)
	cond.Signal()// 3秒之后 下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 1)
	cond.Broadcast()//3秒之后 下发广播给所有等待的goroutine
	time.Sleep(time.Second * 10)
	fmt.Println("finish all")

}