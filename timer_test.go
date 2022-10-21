package demo

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TimerSend(t, content, to string) (tr *time.Timer) {
	var hour, min int

	ts := strings.Split(t, ":")
	if len(ts) != 2 {
		fmt.Errorf("time format split error")
	}

	var err error
	if hour, err = strconv.Atoi(ts[0]); err != nil {
		fmt.Errorf("time format error: %v", err)
	}

	if min, err = strconv.Atoi(ts[1]); err != nil {
		fmt.Errorf("time format error: %v", err)
	}

	if hour < 0 || hour > 23 || min < 0 || min > 59 {
		fmt.Errorf("time format error: %v", t)
	}
	s := time.Now()
	c := time.Date(s.Year(), s.Month(), s.Day(), 0, 0, 0, 0, time.Local)

	if d := time.Duration(hour)*time.Hour + time.Duration(min)*time.Minute - s.Sub(c); d > 0 {
		tr = time.NewTimer(d)
		go func() {
			defer tr.Stop()
			<-tr.C
			fmt.Println(content, " to ", to)
		}()
	}
	return
}

func startTimer(f func()) chan bool {
	done := make(chan bool, 1)
	go func() {
		timer := time.NewTimer(5 * time.Second)
		defer timer.Stop()
		for {
			select {
			case <-timer.C:
				f()
			case <-done:
				return
			}
		}
	}()
	return done
}

func TestTimer(t *testing.T) {
	// 使用方法
	// done := startTimer(func() {
	// 	fmt.Println("Timer 1 expired")
	// })
	// close(done)
	ts := make(map[string]*time.Timer)
	ts["1"] = TimerSend("15:50", "content1", "to1")
	ts["2"] = TimerSend("15:51", "content2", "to2")
	ts["1"].Stop()
	time.Sleep(120 * time.Second)
	/*
		if d := GoAt("11:23"); d > 0 {
			t := time.NewTimer(d)
			go func() {
				<-t.C

			}()
		}
	*/
	d, _ := time.ParseDuration("10s")
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
