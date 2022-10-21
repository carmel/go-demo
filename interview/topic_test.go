package interview

import (
	"fmt"
	"reflect"
	"runtime"
	"sync"
	"testing"
)

// 不使用中间变量交换两个变量的值
func TestAPlusB(t *testing.T) {
	a, b := 10, 20

	a = a + b
	b = a - b
	a = a - b
	t.Error(a, b)

	b, a = a, b
	t.Error(a, b)
}

type student struct {
	Name string
	Age  int
}

func TestStruct(t *testing.T) {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	t.Error(m)
}

func TestStruct1(t *testing.T) {
	kv := map[string]student{"menglu": {Age: 21}}
	// kv["menglu"].Age = 22 此处无法寻址，非常规寻址值
	s := []student{{Age: 21}}
	s[0].Age = 22
	fmt.Println(kv, s)

	a := make([]int, 3)
	fmt.Println(len(a))

	b := make(map[string]string, 3)
	fmt.Println(len(b))
}

const (
	a = iota
	b = iota
)
const (
	name = "menglu"
	c    = iota
	d    = iota
)

// const中每新增一行常量声明将使iota计数一次
// iota可理解为const语句块中的行索引
func TestIota(t *testing.T) {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println([...]string{"1"} == [...]string{"1"})
	fmt.Println(reflect.TypeOf([...]string{"1"}), reflect.ValueOf([...]string{"1"}).Kind())
	fmt.Println(reflect.TypeOf([]string{"1"}), reflect.ValueOf([]string{"1"}).Kind())
	// fmt.Println([]string{"1"} == []string{"1"})
}

func TestWaitgroup(t *testing.T) {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
