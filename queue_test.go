package demo

import (
	"container/list"
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	// 创建队列
	l := list.New()
	// 入队or压栈
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	// 出队
	i1 := l.Front()
	l.Remove(i1)
	fmt.Printf("出队: %d\n", i1.Value)
	// 出栈
	i4 := l.Back()
	l.Remove(i4)
	fmt.Printf("出栈: %d\n", i4.Value)
	fmt.Println("当前size: ", l.Len())

	i1 = l.Front()
	l.Remove(i1)
	i1 = l.Front()
	l.Remove(i1)
}
