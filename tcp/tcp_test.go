package tcp_test

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"sync"
	"testing"
)

const Msg_Header = "12345678"

func Connector(addr string, sendTimes int) {
	const msg = "我是一个完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整完完整整的数据包"
	// 创建连接
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("连接错误：", err)
		return
	}

	// 循环请求
	for i := 0; i < sendTimes; i++ {
		err := Encode(conn, fmt.Sprintf("%d:%s", i, msg))
		if err != nil {
			fmt.Println("发送错误：", err)
			break
		}
	}
}

// 接收器
type Acceptor struct {
	ls            net.Listener   // 保存侦听器
	wg            sync.WaitGroup // 侦听器同步
	OnSeesionData func(net.Conn, []byte) bool
}

// 异步开始侦听
func (a *Acceptor) Start(addr string) {
	go func(addr string) {
		a.wg.Add(1) // 侦听开始时，添加一个任务
		defer a.wg.Done()

		var err error
		a.ls, err = net.Listen("tcp", addr)
		if err != nil {
			fmt.Println("listen err = ", err)
			return
		}
		// 侦听循环
		for {
			conn, err := a.ls.Accept() // 新连接没到来时，Accept阻塞
			if err != nil {
				break
			}
			go handleSession(conn, a.OnSeesionData)
		}
	}(addr)
}

// 停止侦听
func (a *Acceptor) Stop() {
	a.ls.Close()
}

// 等待侦听完全停止
func (a *Acceptor) Wait() {
	a.wg.Wait()
}

func NewAcceptor() *Acceptor {
	return &Acceptor{}
}

func handleSession(conn net.Conn, callback func(net.Conn, []byte) bool) {
	dataReader := bufio.NewReader(conn) // 创建socket读取器
	for {                               // 循环接收数据
		data, err := Decode(dataReader)
		if err != nil || !callback(conn, data) {
			// 回调要求退出
			conn.Close()
			break
		}
	}
}

func Encode(bytesBuffer io.Writer, content string) error {
	//msg_header+content_len+content
	//8			+4			+content
	if err := binary.Write(bytesBuffer, binary.BigEndian, []byte(Msg_Header)); err != nil {
		return err
	}
	clen := int32(len([]byte(content)))
	if err := binary.Write(bytesBuffer, binary.BigEndian, clen); err != nil {
		return err
	}
	if err := binary.Write(bytesBuffer, binary.BigEndian, []byte(content)); err != nil {
		return err
	}
	return nil
}

func Decode(bytesBuffer io.Reader) (bodyBuf []byte, err error) {
	MagicBuf := make([]byte, len(Msg_Header))
	if _, err = io.ReadFull(bytesBuffer, MagicBuf); err != nil {
		return nil, err
	}
	if string(MagicBuf) != Msg_Header {
		return nil, errors.New("msg_header error")
	}

	lengthBuf := make([]byte, 4)
	if _, err = io.ReadFull(bytesBuffer, lengthBuf); err != nil {
		return nil, err
	}

	length := binary.BigEndian.Uint32(lengthBuf)
	bodyBuf = make([]byte, length)
	if _, err = io.ReadFull(bytesBuffer, bodyBuf); err != nil {
		return nil, err
	}
	return bodyBuf, err
}

func TestMain(t *testing.T) {
	const TESTCOUNT = 10000 // 测试次数
	const addr = "127.0.0.1:3000"
	var recCounter int // 接收器

	a := NewAcceptor()
	a.Start(addr)
	a.OnSeesionData = func(conn net.Conn, data []byte) bool {
		str := string(data)
		fmt.Println(str)
		n, err := strconv.Atoi(strings.Split(str, ":")[0])
		if err != nil || recCounter != n {
			panic("failed")
		}
		recCounter++

		if recCounter >= TESTCOUNT {
			a.Stop()
			return false
		}
		return true
	}

	// 连接器不断发送数据
	Connector(addr, TESTCOUNT)

	// 等待侦听器结束
	a.Wait()
}
