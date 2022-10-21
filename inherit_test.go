package demo

import (
	"fmt"
	"testing"
)

type Person struct {
	name string
}

func (p *Person) Say() {
	fmt.Println("My name is ", p.name)
}

type Student struct {
	Person
}

func TestInherit(t *testing.T) {
	var s = Student{Person{"Vector"}}
	fmt.Println(s.name)
	s.Say()
}
