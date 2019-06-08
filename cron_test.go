package test

import (
	"fmt"
	"time"

	"github.com/jasonlvhit/gocron"
)

func task1() {
	fmt.Println(`task 1 excute`)
}

func task2() {
	fmt.Println(`task 2 excute`)
}

func main() {
	m := make(map[string]func())
	m["1"] = task1
	m["2"] = task2

	s := gocron.NewScheduler()
	s.Every(1, false).Second().Do(m["1"])
	s.Every(2, true).Seconds().Do(m["2"])
	go func() {
		time.Sleep(6 * time.Second)
		s.Remove(m["1"])
		time.Sleep(8 * time.Second)
		s.Every(1, true).Second().Do(m["1"])
	}()
	<-s.Start()

}
