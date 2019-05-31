package test

import (
	"fmt"
	"testing"

	"github.com/mitchellh/mapstructure"
)

type User struct {
	Name  string
	Phone string
}

func TestMain(m *testing.M) {

	mp := map[string]string{
		"name":  "vector",
		"phone": "186",
	}
	var user User
	fmt.Println(mapstructure.Decode(mp, &user))
	fmt.Println(user)

}
