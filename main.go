package main

import (
	"log"
	"net"

	pb "github.com/Cyber-Nomad-Collective/PlanPilot_Interface/gen/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	client, err := NewOpenRouterClient()
	if err != nil {
		log.Fatal(err)
	}

	client.Models()

	grpcServer := grpc.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	pb.RegisterGreeterServer(
		grpcServer,
		NewGreeterServer(client),
	)

	log.Println("gRPC server listening on :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
