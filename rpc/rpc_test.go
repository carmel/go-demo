package service_test

import (
	"errors"
	"fmt"
	service "go-demo/rpc"
	"net/rpc"
	"testing"
)

type Device struct{}

func (this *Device) TestRPCMethod(data string, reply *string) error {
	if data == "success" {
		*reply = "your data is " + data
		return nil
	}
	return errors.New("something gonna error")
}

func TestMain(m *testing.M) {
	service.ListenRPC("5514", &Device{})
	m.Run()
}

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
