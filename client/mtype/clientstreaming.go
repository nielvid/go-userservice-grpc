package mtype

import (
	"context"
	"log"
	"time"

	pb "github.com/nielvid/go-userservice-grpc/proto"
)

//streaming data to the server
func VerifyUsers(client pb.UserServiceClient) {

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()

	usersID := []string{"x", "s", "o"}

	stream, err := client.VerifyUsers(context.Background())
	if err != nil {
		log.Fatalf("RPC failed: %v", err)
	}

	for _, id := range usersID {
		req := &pb.Params{
			Id: id,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error sending stream: %v", err)
		}

		log.Printf("user id sent for verification %v", id)
		time.Sleep(3 * time.Second)
	}
	log.Println("finished streaming usersId to server for verifiaction")
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error receiving response: %v", err)
	}
	log.Println(res, "responze received from the server")

}
