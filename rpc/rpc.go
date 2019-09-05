package service

import (
	cfg "flue/transport/conf"
	"fmt"
	"net"
	"net/rpc"
	"time"
)

func ListenRPC(port string, rcvr interface{}) {
	l, e := net.Listen("tcp", ":"+port)
	if e != nil {
		fmt.Println("Error: listen "+port+" error:", e)
	} else {
		fmt.Println("RPC Server listen to " + port)
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
				p.Register(rcvr)
				p.ServeConn(conn)
			}(conn)
		}
	}()
}

func Call(server, rpcname string, args, reply interface{}) error {
	conn, err := net.DialTimeout("tcp", server, time.Duration(cfg.Conf.Server.DailTimeout)*time.Second)
	if err == nil {
		conn.SetDeadline(time.Now().Add(time.Duration(cfg.Conf.Server.DeadLine) * time.Second))
		c := rpc.NewClient(conn)
		// err = c.Call(rpcname, args, reply) // 同步方式
		rpcCall := c.Go(rpcname, args, reply, nil) // 异步方式
		cl := <-rpcCall.Done
		return cl.Error
	}
	return err
}
