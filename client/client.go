package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
	"test_grpc/gateway/config"
	pb "test_grpc/protos"
	"time"
)

var (
	configPath = flag.String("c", "./gateway/config/config.conf", "Application config file")
)

func main() {
	flag.Parse()
	if configPath == nil {
		os.Exit(1)
	}
	config.InitConfig(*configPath, config.Settings)

	// Set up a connection to the server.
	address := fmt.Sprintf("%v:%v", config.Settings.Server.GrpcHost, config.Settings.Server.GrpcPort)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)

	// Contact the server and print out its response.
	name := "haohao_client"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name, Age: 20})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
