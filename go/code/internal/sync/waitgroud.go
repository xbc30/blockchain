package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	seconds := [3]int{1,2,3}

	for i, s := range seconds {
		// 计数加 1
		fmt.Println(i)
		wg.Add(1)
		go func(i, s int) {
			// 计数减 1
			defer wg.Done()
			fmt.Printf("goroutine%d 结束\n", i)
		}(i, s)
	}

	// 等待执行结束
	wg.Wait()
	fmt.Println("所有 goroutine 执行结束")
}