package test

import (
	"flag"
	"fmt"
)

//"bufio"

func main() {

	typ := flag.String("typ", "1", "test")
	flag.Parse()

	defer func() {
		if ok := recover(); ok != nil {
			fmt.Println("recover")
		}
	}()
	switch *typ {
	case "1":
		fmt.Println(`input value is 1`)
	case "2":
		fmt.Println(`input value is 1`)

	default:
		panic(`input value is not 1 or 2`)
	}
}
