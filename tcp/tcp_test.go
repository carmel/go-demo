package tcp_test

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"testing"
)

const Msg_Header = "##"

// 原报文: "##0158QN=20180531225849000;ST=31;CN=2011;PW=123456;MN=010000A8900016F018010001;Flag=5;CP=&&DataTime=20180531225849;a34041-Rtd=0.00,a34041-Flag=N;SB7-RS=2;SB8-RS=2&&9181\r\n"
// 报文格式: 2byte包头(##) + 数据段长度(4byte十进制数) + CRC校验(4byte十六进制数) + 包尾(\r\n，2byte)
func Connector(addr string, sendTimes int) {
	const msg = "QN=20180531225849000;ST=31;CN=2011;PW=123456;MN=010000A8900016F018010001;Flag=5;CP=&&DataTime=20180531225849;a34041-Rtd=0.00,a34041-Flag=N;SB7-RS=2;SB8-RS=2&&"
	// 创建连接
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("连接错误：", err)
		return
	}

	// 循环请求
	for i := 0; i < sendTimes; i++ {
		err := Encode(conn, msg)
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
	binary.Write(bytesBuffer, binary.BigEndian, []byte("9181"))
	binary.Write(bytesBuffer, binary.BigEndian, []byte("\r\n"))
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
	// 读出最后的6个字节
	io.ReadFull(bytesBuffer, make([]byte, 6))
	// bytesBuffer.Read()
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
