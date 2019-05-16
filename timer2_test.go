package main

import (
	"fmt"
	"time"
)

func main() {
	d, err := time.ParseDuration("10s")
	//fmt.Println(d.String(), err)
	//fmt.Println(time.Now().Add(d))

	timer := time.NewTimer(d)
	defer timer.Stop()
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case c := <-ticker.C:
			fmt.Println(`------当前日期:`, c)
		case <-timer.C:
			fmt.Println("-----执行定时任务")
			d, _ = time.ParseDuration("5s")
			timer.Reset(d)
		}
	}
}
