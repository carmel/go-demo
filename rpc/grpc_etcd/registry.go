package balancer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"hash/crc32"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var prefix = "/registry/server/"

type Registry interface {
	RegistryNode(addr string) error
	UnRegistry()
}

type registryServer struct {
	cli        *clientv3.Client
	stop       chan bool
	isRegistry bool
	options    Options
	leaseID    clientv3.LeaseID
}

// type PutNode struct {
// 	Addr string `json:"addr"`
// }

type Node struct {
	Id   uint32 `json:"id"`
	Addr string `json:"addr"`
}

type Options struct {
	name   string
	ttl    int64
	config clientv3.Config
}

func NewRegistry(options Options) (Registry, error) {
	cli, err := clientv3.New(options.config)
	if err != nil {
		return nil, err
	}
	return &registryServer{
		stop:       make(chan bool),
		options:    options,
		isRegistry: false,
		cli:        cli,
	}, nil
}

func (s *registryServer) RegistryNode(addr string) error {
	if s.isRegistry {
		return errors.New("only one node can be registered")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.options.ttl)*time.Second)
	defer cancel()
	grant, err := s.cli.Grant(context.Background(), s.options.ttl)
	if err != nil {
		return err
	}
	var node = Node{
		Id:   s.HashKey(addr),
		Addr: addr,
	}
	nodeVal, err := s.GetVal(node)
	if err != nil {
		return err
	}
	_, err = s.cli.Put(ctx, s.GetKey(node), nodeVal, clientv3.WithLease(grant.ID))
	if err != nil {
		return err
	}
	s.leaseID = grant.ID
	s.isRegistry = true
	go s.KeepAlive()
	return nil
}

func (s *registryServer) UnRegistry() {
	s.stop <- true
}

func (s *registryServer) Revoke() error {
	_, err := s.cli.Revoke(context.TODO(), s.leaseID)
	if err != nil {
		log.Printf("[Revoke] err : %s", err.Error())
	}
	s.isRegistry = false
	return err
}

// Lease提供了几个功能：

// Grant：分配一个租约。
// Revoke：释放一个租约。
// TimeToLive：获取剩余TTL时间。
// Leases：列举所有etcd中的租约。
// KeepAlive：自动定时的续约某个租约。
// KeepAliveOnce：为某个租约续约一次。
// Close：貌似是关闭当前客户端建立的所有租约。

func (s *registryServer) KeepAlive() error {
	keepAliveCh, err := s.cli.KeepAlive(context.TODO(), s.leaseID)
	if err != nil {
		log.Printf("[KeepAlive] err : %s", err.Error())
		return err
	}
	for {
		select {
		case <-s.stop:
			_ = s.Revoke()
			return nil
		case _, ok := <-keepAliveCh:
			if !ok {
				_ = s.Revoke()
				return nil
			}
		}
	}
}

func (s *registryServer) GetKey(node Node) string {
	return fmt.Sprintf("%s%s/%d", prefix, s.options.name, s.HashKey(node.Addr))
}

func (s *registryServer) GetVal(node Node) (string, error) {
	data, err := json.Marshal(&node)
	return string(data), err
}

func (e *registryServer) HashKey(addr string) uint32 {
	return crc32.ChecksumIEEE([]byte(addr))
}
