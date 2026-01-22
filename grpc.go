package main

import (
	"context"

	pb "github.com/Cyber-Nomad-Collective/PlanPilot_Interface/gen/proto"
)

type GreeterServer struct {
	openRouterClient OpenRouterClient
	pb.UnimplementedGreeterServer
}

func NewGreeterServer(openRouterClient *OpenRouterClient) *GreeterServer {
	return &GreeterServer{
		openRouterClient: *openRouterClient,
	}
}

func (s *GreeterServer) SayHello(
	ctx context.Context,
	req *pb.HelloRequest,
) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
