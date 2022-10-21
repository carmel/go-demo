package demo

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

type U struct {
	Id    int
	Name  string
	Birth time.Time
}

func (u U) Say() {
	fmt.Println("Hello, World!")
}

func StructInfo(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type of o is ", t.Name())
	v := reflect.ValueOf(o)
	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}
}

/*
Type of o is  User
Fields:
    Id: int = 12
  Name: string = Chen
 Birth: time.Time = 2018-01-11 20:50:58.8164074 +0800 CST m=+0.002000101
*/

func TestReflect(ts *testing.T) {
	//	u := User{12, "Chen", time.Now()}
	//	StructInfo(u)
	t := reflect.New(reflect.TypeOf("U"))
	fmt.Println(t)
}
