package main
import (
	"fmt"
	"sync"
)
var x  = 0
func increment(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<- ch
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, ch)
	}
	w.Wait()
	// 有缓冲的channel，如果在 go routine 中使用，一定要做适当的延时，否则会输出来不及，因为main已经跑完了，所以延时一会，等待 go routine
	// time.Sleep(time.Second * 1)
	fmt.Println("final value of x", x)
}
