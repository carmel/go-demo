package tests_test

import (
	"fmt"
	"go-demo/tests"
	"log"
	"testing"

	"github.com/stretchr/testify/mock"
)

// 一个 "Production" 例子
func TestMain(m *testing.M) {
	/*
		fmt.Println("Hello World")
		smsService := tests.SMSService{}
		myService := tests.MyService{smsService}
		myService.ChargeCustomer(100)
		m.Run()*/
	d := tests.NewDB()

	g := tests.NewGreeter(d, "en")
	log.Println(g.Greet())             // Message is: hello
	log.Println(g.GreetInDefaultMsg()) // Message is: default message

	g = tests.NewGreeter(d, "es")
	log.Println(g.Greet()) // Message is: holla

	g = tests.NewGreeter(d, "random")
	log.Println(g.Greet()) // Message is: bzzzz
}

// smsServiceMock
type smsServiceMock struct {
	mock.Mock
}

// 我们模拟的 smsService 方法
func (m *smsServiceMock) SendChargeNotification(value int) bool {
	fmt.Println("Mocked charge notification function")
	fmt.Printf("Value passed in: %d\n", value)
	// 这将记录方法被调用以及被调用时传进来的参数值
	args := m.Called(value)
	// 它将返回任何我需要返回的
	// 这种情况下模拟一个 SMS Service Notification 被发送出去
	return args.Bool(0)
}

// 我们将实现 MessageService 接口
// 这就意味着我们不得不改写在接口中定义的所有方法
func (m *smsServiceMock) DummyFunc() {
	fmt.Println("Dummy")
}

// TestChargeCustomer 是个奇迹发生的地方
// 在这里我们将创建 SMSService mock
func TestChargeCustomer(t *testing.T) {
	smsService := new(smsServiceMock)

	// 然后我们将定义当 100 传递给 SendChargeNotification 时，需要返回什么
	// 在这里，我们希望它在成功发送通知后返回 true
	smsService.On("SendChargeNotification", 1000).Return(true, nil)

	// 接下来，我们要定义要测试的服务
	myService := tests.MyService{smsService}
	// 然后调用方法
	myService.ChargeCustomer(100)

	// 最后，我们验证 myService.ChargeCustomer 调用了我们模拟的 SendChargeNotification 方法
	smsService.AssertExpectations(t)
}
