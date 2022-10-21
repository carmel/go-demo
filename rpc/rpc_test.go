package rpc

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"testing"
	"time"
)

type Device struct{}

func (this *Device) TestRPCMethod(data string, reply *string) error {
	if data == "success" {
		*reply = "your data is " + data
		return nil
	}
	return errors.New("something gonna error")
}

// func TestMain(m *testing.M) {
// 	service.ListenRPC("5514", &Device{})
// 	m.Run()
// }

func TestRPC(t *testing.T) {
	clientTCP, _ := rpc.Dial("tcp", ":5514")
	quotient := new(string)
	divCall := clientTCP.Go("Device.TestRPCMethod", "success", quotient, nil)
	replyCall := <-divCall.Done

	if replyCall.Error != nil {
		fmt.Println(replyCall.Error)
	} else {
		fmt.Println("success test: ", *quotient)
	}

	divCall = clientTCP.Go("Device.TestRPCMethod", "fail", quotient, nil)
	replyCall = <-divCall.Done

	if replyCall.Error != nil {
		fmt.Println("fail test: ", replyCall.Error)
	} else {
		fmt.Println(*quotient)
	}

}

func TestServer(t *testing.T) {
	l, e := net.Listen("tcp", ":8080")
	if e != nil {
		fmt.Println("Error: listen 8080 error:", e)
	} else {
		fmt.Println("RPC Server listen to 8080")
	}
	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				fmt.Println("Error: accept rpc connection", err.Error())
				continue
			}
			go func(conn net.Conn) {
				defer conn.Close()
				p := rpc.NewServer()
				p.Register(new(Device))
				p.ServeConn(conn)
			}(conn)
		}
	}()
}

func TestRpcClient(t *testing.T) {
	conn, err := net.DialTimeout("tcp", ":8080", time.Duration(120)*time.Second)
	if err == nil {
		conn.SetDeadline(time.Now().Add(time.Duration(120) * time.Second))
		c := rpc.NewClient(conn)
		// err = c.Call(rpcname, args, reply) // 同步方式
		var reply string
		rpcCall := c.Go("Device.TestRPCMethod", "message", reply, nil) // 异步方式
		<-rpcCall.Done
	}
}
