package balancer

import (
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testing"
	"time"

	"fmt"
	pb "go-demo/rpc/helloworld"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/pkg/transport"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	DialTimeout = 20 * time.Second
	// requestTimeout = 10 * time.Second
)

// 定义helloService并实现约定的接口
type helloService struct{}

// SayHello 实现Hello服务接口
func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	resp := new(pb.HelloReply)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)

	return resp, nil
}

// reg=122.51.240.88:2379
func register(port, reg string, op Options) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}
	defer lis.Close()

	s := grpc.NewServer()
	defer s.GracefulStop()

	pb.RegisterGreeterServer(s, &helloService{})

	go func() {
		r, err := NewRegistry(op)
		if err != nil {
			log.Fatal(err)
			return
		}
		err = r.RegistryNode(fmt.Sprintf("127.0.0.1:%s", port))
		if err != nil {
			log.Fatal(err)
			return
		}

		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
		go func() {
			s := <-ch
			r.UnRegistry()
			if i, ok := s.(syscall.Signal); ok {
				os.Exit(int(i))
			} else {
				os.Exit(0)
			}

		}()
	}()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func TestClient(t *testing.T) {
	tlsInfo := transport.TLSInfo{
		CertFile:      `../ca/client.pem`,
		KeyFile:       `../ca/client.key`,
		TrustedCAFile: `../ca/root.pem`,
	}
	// var config *tls.Config
	config, err := tlsInfo.ClientConfig()
	if err != nil {
		panic(err)
	}

	var op = Options{
		name: "svc.info",
		ttl:  10,
		config: clientv3.Config{
			Endpoints:   []string{"122.51.240.88:2379"},
			DialTimeout: DialTimeout,
			TLS:         config,
		},
	}
	for i := 1; i <= 3; i++ {
		go register(fmt.Sprintf("%d%d%d%d", i, i, i, i), "122.51.240.88:2379", op)
		// if i == 3 {
		// 	go func() {
		// 		time.Sleep(time.Second * 20)
		// 		r.UnRegistry()
		// 	}()
		// }
	}

	var sop = SelectorOptions{
		name: "svc.info",
		config: clientv3.Config{
			Endpoints:   []string{"122.51.240.88:2379"},
			DialTimeout: DialTimeout,
			TLS:         config,
		},
	}

	s, err := NewSelector(sop)
	if err != nil {
		t.Error(err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for c := range ticker.C {
			val, err := s.Next()
			if err != nil {
				t.Error(err)
				continue
			}

			conn, err := grpc.Dial(val.Addr, grpc.WithInsecure(), grpc.EmptyDialOption{})
			if err != nil {
				panic(err)
			}
			client := pb.NewGreeterClient(conn)
			resp, err := client.SayHello(
				context.Background(),
				&pb.HelloRequest{Name: fmt.Sprintf("server: %s %ds", val.Addr, c.Second())},
				grpc.WaitForReady(true),
			)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(resp)
			}
		}
		wg.Done()
	}()

	for i := 4; i <= 6; i++ {
		go register(fmt.Sprintf("%d%d%d%d", i, i, i, i), "122.51.240.88:2379", op)
		// if i == 3 {
		// 	go func() {
		// 		time.Sleep(time.Second * 20)
		// 		r.UnRegistry()
		// 	}()
		// }
	}
	wg.Wait()
}
