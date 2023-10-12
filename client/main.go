package main

import (
	"log"

	"github.com/nielvid/go-userservice-grpc/client/mtype"
	pb "github.com/nielvid/go-userservice-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverUrl = "localhost:5000"
	connError  = "error connecting to server"
)

func main() {

	conn, err := grpc.Dial(serverUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf(connError, err)
	}

	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	mtype.CreateUser(client)
	// mtype.FetchUsers(client)
	// mtype.Chat(client)

}
