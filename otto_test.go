package demo

import (
	"fmt"
	"testing"

	"github.com/robertkrimen/otto"
)

func TestOtto(t *testing.T) {
	schema := `
	var d = document.getElementById("test");
	d.innerHTML = "Hello, World!";
	`
	vm := otto.New()
	vm.Set("self", "dfdfdfd")
	val, err := vm.Eval(schema)
	if err != nil {
		fmt.Println(err)
	}
	val.Object()
	fmt.Println(val)
}
