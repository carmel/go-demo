package demo

import (
	"sync"
)

type Singleton struct{}

var instance *Singleton

func GetInstance() *Singleton {
	var once sync.Once
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}
