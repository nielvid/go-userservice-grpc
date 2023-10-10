package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	pb "github.com/nielvid/go-userservice-grpc/proto"
	"github.com/nielvid/go-userservice-grpc/users"
	"github.com/nielvid/go-userservice-grpc/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	godotenv.Load()
	utils.ValidateEnvVars()
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8090"
	}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &users.UserServer{})
	reflection.Register(server)
	fmt.Printf("Server listening on %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)

	}
}
