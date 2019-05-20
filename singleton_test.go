package test

import (
	"sync"
)

type Singleton struct{}

var instance *Singelton

func GetInstance() *Singleton {
	sync.Once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}
