package test

import (
	"fmt"
	"github.com/robertkrimen/otto"
)

func main() {
	schema := `
	var d = document.getElementById("test");
	d.innerHTML = "Hello, World!";
	`
	vm := otto.New()
	vm.Set("self", "dfdfdfd")
	val, err := vm.Eval(schema)
	if err != nil {
		fmt.Println("%v\n", err)
	}
	val.Object()
	fmt.Println(val)
}
