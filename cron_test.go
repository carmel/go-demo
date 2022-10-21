package demo

import (
	"fmt"
	"testing"

	"github.com/robfig/cron/v3"
)

func task1() {
	fmt.Println(`task 1 excute`)
}

func task2() {
	fmt.Println(`task 2 excute`)
}

func TestCron(t *testing.T) {
	m := make(map[string]func())
	m["1"] = task1
	m["2"] = task2

	cn := cron.New()
	cn.AddFunc("CRON_TZ=Asia/Shanghai 0 0 * * ?", func() {
		t.Logf(`[zero time cron]`)
	})
	cn.Start()

}
