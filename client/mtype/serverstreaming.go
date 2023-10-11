package mtype

import (
	
	"context"
	"io"
	"log"

	pb "github.com/nielvid/go-userservice-grpc/proto"
) 


//recieving stream of data from the server
func FetchUsers(client pb.UserServiceClient) {
	

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()

	req := &pb.NoParams{}

	userStream, err := client.FetchUsers(context.Background(), req)
	if err != nil {
		log.Fatalf("RPC failed: %v", err)
	}
	log.Println("streaming users from database")
	for {
		user, err := userStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error receiving: %v", err)
		}
		log.Println(user, "user from db")
	}
	log.Println("finished streaming users from database")

}
