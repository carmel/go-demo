package main

import (
	"fmt"
)

func main() {
	fmt.Println("----循环扫码---")
	var s string
	str := make(chan string, 1)
	go func() {
		for {
			fmt.Scan(&s)
			str <- s
		}
	}()

	for {
		fmt.Println("hello,", <-str)
	}

}
