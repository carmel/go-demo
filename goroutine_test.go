package demo

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

var (
	chanNum   = 3                     //启动的数量
	readChan  = make(chan int)        //操作信息的channel
	limitChan = make(chan bool, 1000) //限制goroutine数量的channel，此处限制1000个
)

//初始人方法
func init() {
	fmt.Println("init")

	for i := 0; i < chanNum; i++ {
		go Queue(i, readChan) //开启工作池
	}
}

func TestGoroutine(t *testing.T) {
	fmt.Println("main")

	//启一个go方法 ， 无限制的往readChan里塞数据
	go func() {
		for {
			readChan <- 1
		}
	}()

	//监听到键盘事件后程序退出
	var input string
	fmt.Scanln(&input)
}

//工作池
func Queue(qid int, rchan chan int) {
	var dat int
	t := time.Tick(time.Second) //定时器，一秒
	for {
		select {
		case d := <-rchan:
			limitChan <- true //缓冲区满之后阻塞，后面的readChan将等待
			dat += d
			go worker(qid, dat) //每从channel接到一个数据就起一个goroutine,limitChan会限制goroutine的数量
		case <-t:
			showGoNum(qid) //定时器，每秒打印一次当前 goroutine数量
		}
	}
}

func worker(qid, i int) {
	//处理成功后在限制goroutine的Channel缓冲区里取一个数据，limitChan就可以再写入
	//使用Defer 确保limitChan的一个缓冲区被释放
	defer func() {
		<-limitChan
	}()
	time.Sleep(time.Millisecond * 100) //模拟程序处理耗时

	fmt.Println(qid, "===========", i)
}

//显示当前goroutine数量
func showGoNum(qid int) {
	fmt.Printf("%d====numGo:==%d\n", qid, runtime.NumGoroutine())
}
