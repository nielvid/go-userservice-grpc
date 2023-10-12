package mtype

import (
	"context"
	"log"

	pb "github.com/nielvid/go-userservice-grpc/proto"
)

// recieving stream of data from the server
func CreateUser(client pb.UserServiceClient) {

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	var Phone = "0806733"
	req := &pb.UserParams{Firstname: "ade", Lastname: "niel", Email: "XXXXXXXXXXXXX", Password: "123", PhoneNumber: &Phone}

	res, err := client.CreateUser(context.Background(), req)
	if err != nil {
		log.Fatalf("RPC failed: %v", err)
	}
	log.Println("streaming users from database")

	log.Println("response", res)

}
