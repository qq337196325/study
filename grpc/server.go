package grpc

import (
	"context"
	"log"
	"net"

	pb "study/proto"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type Server struct {
	pb.UnimplementedGreeterServer
}

func (s *Server) NewGrpcServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpc := grpc.NewServer()
	pb.RegisterGreeterServer(grpc, &Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := grpc.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
