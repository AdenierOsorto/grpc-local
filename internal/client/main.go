package main

import (
	"context"
	"log"

	pb "github.com/AdenierOsorto/grpc-local/internal/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var address string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewUserServiceClient(conn)

	DoUser(c)
}

func DoUser(c pb.UserServiceClient) {
	log.Print("doUser was invoked")
	member := &pb.UpdatePrimaryMemberDTO{
		Title:     "Mr.",
		Gender:    "male",
		FirstName: "Adenier",
		LastName:  "Osorto",
		Email:     "elmer.osorto@uvltd.tech",
		ShirtSize: "xl",
	}
	res, err := c.UpdatePrimaryMember(context.Background(), &pb.UpdatePrimaryMemberRequest{
		Member: member,
	})

	if err != nil {
		log.Fatalf("could not user: %v\n", err)
	}

	log.Printf("User updated: %s\n", res.Message)
}
