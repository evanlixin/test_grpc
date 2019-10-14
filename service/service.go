package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"test_grpc/protos"
	pb "test_grpc/protos"
)

type HelloService struct {
}

//	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReponse, error)
func (s *HelloService) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloReponse, error) {
	log.Printf("Received: %v", in.Name)
	return &hello.HelloReponse{Reply: in.Name}, nil
}
func main() {
	c, err := credentials.NewServerTLSFromFile("../certs/server/server.pem", "../certs/server/server.key")
	if err != nil {
		log.Fatalf("credentials.NewServerTLSFromFile err: %v", err)
	}
	server := grpc.NewServer(grpc.Creds(c))
	//server := grpc.NewServer()

	pb.RegisterHelloServiceServer(server, &HelloService{})

	lis, err := net.Listen("tcp", ":50123")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
