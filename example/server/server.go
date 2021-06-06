package server

import (
	"context"
	"fmt"
	pb "github.com/romannikov/example/api"
	"github.com/romannikov/example/config"
	"google.golang.org/grpc"
	"net"
)

type exampleServer struct {
	pb.UnimplementedExampleServer
	config *config.ExampleConfig
}

func (s *exampleServer) Exam(cxt context.Context, request *pb.ExampleRequest) (*pb.ExampleResponse, error) {
	return &pb.ExampleResponse{Result: 0}, nil
}

func Start(cfg *config.ExampleConfig) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.DbHost, cfg.DbPort))
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	pb.RegisterExampleServer(grpcServer, &exampleServer{pb.UnimplementedExampleServer{}, cfg})
	return grpcServer.Serve(lis)
}
