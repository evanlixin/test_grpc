package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"test_grpc/gateway/config"
	"test_grpc/protos"
	pb "test_grpc/protos"
)

var (
	configPath = flag.String("c", "./gateway/config/config.conf", "Application config file")
)

type HelloService struct {
}

//	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
func (s *HelloService) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	var messge string = fmt.Sprintf("Hello , %v ; your age is %d", in.Name, in.Age)
	return &hello.HelloReply{Message: messge}, nil
}
func main() {

	flag.Parse()
	if configPath == nil {
		os.Exit(1)
	}
	config.InitConfig(*configPath, config.Settings)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", config.Settings.Server.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcService := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcService, &HelloService{})
	if err := grpcService.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
