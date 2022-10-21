package pool

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Task func() error

type Pool struct {
	Tasks       []Task
	concurrency int
	Ch          chan Task
	lock        sync.Mutex
}

func NewPool(concurrency int) *Pool {
	return &Pool{
		Tasks:       []Task{},
		concurrency: concurrency,
		Ch:          make(chan Task, concurrency),
	}
}

func (p *Pool) Stop() {
	p.Tasks = nil
	close(p.Ch)
}

func (p *Pool) AddTask(t Task) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.Tasks = append(p.Tasks, t)
}

func (p *Pool) Start() {
	p.Ch <- nil
	go func() {
		for f := range p.Ch {
			if len(p.Tasks) != 0 {
				p.Ch <- p.Tasks[0]
				p.Tasks = p.Tasks[1:]
			}
			if f != nil {
				go func() {
					if err := f(); err != nil {
						fmt.Println(err)
					}
				}()
			}
		}
	}()
}

func TestFunc(t *testing.T) {
	f := func(id int) Task {
		return func() error {
			fmt.Println(id)
			return nil
		}
	}
	p := NewPool(5)
	defer p.Stop()

	for i := 1; i <= 10000; i++ {
		p.AddTask(f(i))
	}
	p.Start()
	time.Sleep(time.Duration(2) * time.Minute)
}
