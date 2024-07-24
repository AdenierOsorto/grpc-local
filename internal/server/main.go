package main

import (
	"context"
	"log"
	"net"

	pb "github.com/AdenierOsorto/grpc-local/internal/grpc/proto"
	"google.golang.org/grpc"
)

var address string = "0.0.0.0:50051"

type Server struct {
	pb.UserServiceServer
}

func main() {
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("failded to listen on %v\n", err)
	}
	log.Printf("listening on %s\n", address)

	s := grpc.NewServer()

	pb.RegisterUserServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v \n", err)
	}
}

func (s *Server) UpdatePrimaryMember(ctx context.Context, in *pb.UpdatePrimaryMemberRequest) (*pb.UpdatePrimaryMemberResponse, error) {
	log.Printf("Greet functopm was onkoved with %v\n", in)
	return &pb.UpdatePrimaryMemberResponse{
		Success: true,
		Message: "all okay " + in.Member.FirstName,
	}, nil
}
