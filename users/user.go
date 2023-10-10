package users

import (
	"context"
	"log"

	"github.com/nielvid/go-userservice-grpc/database"
	"github.com/nielvid/go-userservice-grpc/models"
	pb "github.com/nielvid/go-userservice-grpc/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserServer struct {
	pb.UserServiceServer
}

var db = database.Connection()

func (s *UserServer) CreateUser(ctx context.Context, req *pb.UserParams) (*pb.User, error) {

	user := &models.User{ID: primitive.NewObjectID(), FirstName: req.Firstname, LastName: req.Lastname, PhoneNumber: *req.PhoneNumber, Email: req.Email, Password: req.Password}
	result, err := db.CreateUser(user)
	if err != nil {
		log.Fatalf("cannot create user :%v", err)
	}
	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Fatalf("failed to cast ObjectID")
	}
	return &pb.User{
		Id:          id.Hex(),
		Firstname:   req.Firstname,
		Lastname:    req.Lastname,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Password:    req.Password,
	}, nil
}
