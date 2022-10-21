package demo

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"testing"
)

func BytesJoin(pBytes ...[]byte) []byte {
	len := len(pBytes)
	s := make([][]byte, len)
	for index := 0; index < len; index++ {
		s[index] = pBytes[index]
	}
	return bytes.Join(s, []byte(""))
}

func TestTcp(ts *testing.T) {
	server, err := net.ResolveTCPAddr("tcp", "challenge.yuansuan.cn:7042")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	conn, err := net.DialTCP("tcp", nil, server)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)
	msg, err := reader.ReadString('\n')
	fmt.Println(msg)
	if err != nil {
		os.Exit(1)
	}
	str := strings.Split(strings.TrimSpace(msg), ":")
	conn.Write([]byte("IAM:" + str[len(str)-1] + ":czk.xu@qq.com\n"))
	msg, err = reader.ReadString('\n')
	var b []byte
	var t []byte
	var c uint32 = 0
	var l uint32 = 0
	var d bytes.Buffer = bytes.Buffer{}
	fmt.Println(msg)
	for {
		msg, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println(d.String())
			ioutil.WriteFile("result", d.Bytes(), 0644)
			os.Exit(1)
		}
		t = []byte(strings.TrimSpace(msg))
		b = BytesJoin(b, t)
		c = uint32(len(b))
		if c > 12 {
			// 大端序转换得到数据长度
			l = binary.BigEndian.Uint32(b[8:12]) + 12
			if c > l {
				//ioutil.WriteFile(string(index)+".png", b[12:l], 0644)
				d.Write(b[12:l])
				b = b[l:]
			}
		}
	}
}
