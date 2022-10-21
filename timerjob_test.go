package demo

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

// t为具体时刻,如"00:00"; d为时间间隔,如"24h"("1d13h12m10s")
// 意即每天晚上凌晨执行一次
func timerJob(t, d string) {
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
		fmt.Errorf("time format error: %v", err)
	}
	dp, err := time.ParseDuration(d)
	if err == nil {
		s := time.Now()
		c := time.Date(s.Year(), s.Month(), s.Day(), 0, 0, 0, 0, time.Local)
		dt := time.Duration(hour)*time.Hour + time.Duration(min)*time.Minute - s.Sub(c)
		if dt < 0 {
			dt += dp
		}
		timer := time.NewTimer(dt)
		defer timer.Stop()
		go func() {
			for {
				select {
				case <-timer.C:
					fmt.Println("-----执行定时任务")
					timer.Reset(dp)
				}
			}
		}()
	} else {
		fmt.Errorf("duration format error")
	}
}

func TestTimeJob(t *testing.T) {
	// 15:13开始运行, 每隔10s运行一次
	timerJob("15:13", "10s")
	ticker := time.NewTicker(2 * time.Second)
	for c := range ticker.C {
		defer ticker.Stop()
		fmt.Println(`------当前时间:`, c)
	}
}
