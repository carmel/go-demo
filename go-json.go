package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type configuration struct {
	Enabled bool
	Path    string
}

func main() {
	file, _ := os.Open("E:/go_project/go-learn-demo/config.json")
	defer file.Close()

	decoder := json.NewDecoder(file)

	conf := configuration{}
	err := decoder.Decode(&conf)
	file.Close()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(conf.Path)
}
