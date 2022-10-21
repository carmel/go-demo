package rpc

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"testing"
	"time"

	pb "go-demo/rpc/helloworld"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func GrpcServer() {
	// creds, err := credentials.NewServerTLSFromFile("ca/server.pem", "ca/server.key")
	// load server cert/key, cacert
	srvcert, err := tls.LoadX509KeyPair("ca/server.pem", "ca/server.key")
	if err != nil {
		log.Fatalf("SERVER: unable to read server key pair: %v", err)
	}

	certPool := x509.NewCertPool()

	caCrt, err := ioutil.ReadFile("ca/root.pem")
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	certPool.AppendCertsFromPEM(caCrt)

	ta := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{srvcert},
		ClientCAs:    certPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	})

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("SERVER: unable to listen: %v", err)
	}
	s := grpc.NewServer(grpc.Creds(ta))

	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func TestGrpcClient(t *testing.T) {

	go GrpcServer()

	// load client cert/key, cacert
	clcert, err := tls.LoadX509KeyPair("ca/client.pem", "ca/client.key")
	if err != nil {
		log.Fatalf("CLIENT: unable to load client pem: %v", err)
	}
	certPool := x509.NewCertPool()

	caCrt, err := ioutil.ReadFile("ca/root.pem")
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	certPool.AppendCertsFromPEM(caCrt)

	ta := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{clcert},
		RootCAs:      certPool,
	})

	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(ta))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
