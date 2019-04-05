package main

import (
	"fmt"
	"runtime"
	"time"
)

var (
	limitChan = make(chan bool, 10) //限制goroutine数量的channel，此处限制1000个
)

func main() {
	t := time.NewTicker(time.Second) //定时器，一秒
	go func() {
		for _ = range t.C {
			fmt.Printf("当前goroutine数目: %d\n", runtime.NumGoroutine())
		}
	}()
	for n := 0; n < 10000; n++ {
		limitChan <- true
		go worker(n)
	}
}

func worker(i int) {
	//处理成功后在限制goroutine的Channel缓冲区里取一个数据，limitChan就可以再写入
	//使用Defer 确保limitChan的一个缓冲区被释放
	defer func() {
		<-limitChan
	}()
	time.Sleep(time.Millisecond * 100) //模拟程序处理耗时
	fmt.Println("序号===========", i)
}
