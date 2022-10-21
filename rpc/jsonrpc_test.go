package rpc

import (
	// "bilibili/library/net/rpc"

	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"
)

type (
	User struct {
		ID   int
		Name string
	}

	UserService struct {
		ID   int
		Name string
	}
)

func (us UserService) GetUser(id int, result *User) error {
	result.ID = 100
	result.Name = "Vector"
	return nil
}

func startServer() {
	us := new(UserService)
	rpc.Register(us)
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}

func TestJsonRpcClient(t *testing.T) {

	go startServer()
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)

	user := User{}

	if err = client.Call("UserService.GetUser", 100, &user); err != nil {
		panic(err)
	}

	log.Println(user)
}
