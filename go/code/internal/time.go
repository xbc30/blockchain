package internal

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 初始化常用 time.Parse() time.Date() time.Now()
	time.Date(2018, 1, 2, 15, 30, 10, 0, time.Local)

	// 格式化输出字符串
	time.Now().Format("2006-01-02 15:04:05") // 2020-04-07 10:11:20

	// 计算当前时区的时间戳
	time.Local = time.FixedZone("CST", 3600*8)
	timestamp := time.Now().Local().Unix(); // time.Now() 可以由任一具体时间time.Parse()替换
	fmt.Println(timestamp)

	// 时间段 Duartion 类型
	tp, _ := time.ParseDuration("1.5s")
	fmt.Println(tp.Truncate(1000), tp.Seconds(), tp.Nanoseconds())

	// 时间运算
	// func Sleep(d Duration)   休眠多少时间，休眠时处于阻塞状态，后续程序无法执行
	time.Sleep(time.Duration(10) * time.Second)
	// func After(d Duration) <-chan Time  非阻塞,可用于延迟
	time.After(time.Duration(10) * time.Second)

	// 加减运算
	fmt.Println(time.Now().Add(time.Duration(10) * time.Second))

	// 前后比较
	dt := time.Date(2018, 1, 10, 0, 0, 1, 100, time.Local)
	fmt.Println(time.Now().After(dt))     // true

	var wg sync.WaitGroup
	wg.Add(2)
	// Ticker类型(心跳时间任务) time.NewTicker定时触发执行任务，当下一次执行到来而当前任务还没有执行结束时，会等待当前任务执行完毕后再执行下一次任务
	// 可通过调用ticker.Stop取消
	ticker := time.NewTicker(1 * time.Minute)
	go func(t *time.Ticker) {
		defer wg.Done()
		for {
			<-t.C
			fmt.Println("get ticker1", time.Now().Format("2006-01-02 15:04:05"))
		}
	}(ticker)

	// Time类型
	timer1 := time.NewTimer(2 * time.Second)
	go func(t *time.Timer) {
		defer wg.Done()
		for {
			<-t.C
			fmt.Println("get timer", time.Now().Format("2006-01-02 15:04:05"))
			//Reset 使 t 重新开始计时，（本方法返回后再）等待时间段 d 过去后到期。如果调用时t
			//还在等待中会返回真；如果 t已经到期或者被停止了会返回假。
			t.Reset(2 * time.Second)
		}
	}(timer1)

	wg.Wait()
}