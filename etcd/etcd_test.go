package etcd

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"testing"
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/pkg/transport"
)

var (
	client         *clientv3.Client
	DialTimeout    = 2 * time.Second
	requestTimeout = 10 * time.Second
)

// 注册到etcd-server
func TestMain(m *testing.M) {
	var exitCode int
	defer os.Exit(exitCode)
	tlsInfo := transport.TLSInfo{
		CertFile:      `ca/client.pem`,
		KeyFile:       `ca/client.key`,
		TrustedCAFile: `ca/root.pem`,
	}
	// var config *tls.Config
	config, err := tlsInfo.ClientConfig()
	if err != nil {
		panic(err)
	}

	client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"122.51.240.88:2379"},
		DialTimeout: DialTimeout,
		TLS:         config,
	})

	if err != nil {
		panic(err)
	}

	exitCode = m.Run()
}

// KV操作
func TestKV(t *testing.T) {

	defer client.Close()

	//控制超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//1. 增-存值
	_, err := client.Put(ctx, "/demo/demo1_key", "demo1_value")
	//操作完毕，cancel掉
	cancel()
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}

	//2. 查-获取值， 也设置超时
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := client.Get(ctx, "/demo/demo1_key")
	// Get查询还可以增加WithPrefix选项，获取某个目录下的所有子元素
	//eg: resp, err := client.Get(ctx, "/demo/", clientv3.WithPrefix())
	cancel()
	if err != nil {
		fmt.Println("get failed err:", err)
		return
	}

	for _, item := range resp.Kvs { //Kvs 返回key的列表
		fmt.Printf("%s : %s \n", item.Key, item.Value)
	}

	//3. 改-修改值
	ctx, _ = context.WithTimeout(context.Background(), time.Second)
	if resp, err := client.Put(ctx, "/demo/demo1_key", "update_value", clientv3.WithPrevKV()); err != nil {
		fmt.Println("get failed err: ", err)
	} else {
		fmt.Println(string(resp.PrevKv.Value))
	}

	//4. 删-删除值
	ctx, _ = context.WithTimeout(context.Background(), time.Second)
	if resp, err := client.Delete(ctx, "/demo/demo1_key"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp.PrevKvs)
	}

	time.Sleep(10 * time.Second)
}

func TestWatchAgent(t *testing.T) {
	defer client.Close()
	rch := client.Watch(context.Background(), "/demo", clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			log.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}

// 事务
func TestTransaction(t *testing.T) {
	defer client.Close()

	var w sync.WaitGroup
	w.Add(10)
	key10 := "setnx"
	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(5 * time.Millisecond)
			//通过key的Create_Revision 是否为 0 来判断key是否存在。其中If，Then 以及 Else 分支都可以包含多个操作。
			//返回的数据包含一个successed字段，当为 true 时代表 If 为真
			_, err := client.Txn(context.Background()).
				If(clientv3.Compare(clientv3.CreateRevision(key10), "=", 0)).
				Then(clientv3.OpPut(key10, fmt.Sprintf("%d", i))).
				Commit()
			if err != nil {
				fmt.Println(err)
			}

			w.Done()
		}(i)
	}
	w.Wait()

	if resp, err := client.Get(context.TODO(), key10); err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
	}
}

// 租期——类似于Redis中的expire (过期时间)
func TestLease(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), requestTimeout)

	kv := clientv3.NewKV(client)
	//Delete all keys
	kv.Delete(ctx, "/demo/demo1_key", clientv3.WithPrefix())

	gr, _ := kv.Get(ctx, "/demo/demo1_key")
	if len(gr.Kvs) == 0 {
		fmt.Println("no key")
	}

	lease, err := client.Grant(ctx, 3)
	if err != nil {
		log.Fatal(err)
	}

	//Insert key with a lease of 3 second TTL
	kv.Put(ctx, "/demo/demo1_key", "demo1_value", clientv3.WithLease(lease.ID))

	gr, _ = kv.Get(ctx, "/demo/demo1_key")
	if len(gr.Kvs) == 1 {
		fmt.Println("Found key")
	}

	//let the TTL expire
	time.Sleep(3 * time.Second)

	gr, _ = kv.Get(ctx, "/demo/demo1_key")
	if len(gr.Kvs) == 0 {
		fmt.Println("no more key")
	}
}
