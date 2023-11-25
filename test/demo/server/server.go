package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/changqing98/idl-hub/idl_gen/demo/idl/demo"
	"google.golang.org/grpc"
	"log"
	"net"
)

type DemoServer struct {
	demo.UnimplementedGreeterServer
}

func (d DemoServer) SayHello(ctx context.Context, request *demo.HelloRequest) (*demo.HelloReply, error) {
	return &demo.HelloReply{
		Message: "hello world",
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8888))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	demo.RegisterGreeterServer(s, &DemoServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
