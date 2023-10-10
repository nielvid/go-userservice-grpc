package users

import (
	"context"
	"log"

	"github.com/nielvid/go-userservice-grpc/auth"
	"github.com/nielvid/go-userservice-grpc/database"
	"github.com/nielvid/go-userservice-grpc/models"
	pb "github.com/nielvid/go-userservice-grpc/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserServer struct {
	pb.UserServiceServer
}


	var (
		access auth.PasetoMaker
		db = database.Connection()
	)

func (s *UserServer) CreateUser(ctx context.Context, req *pb.UserParams) (*pb.AuthUser, error) {

	user := &models.User{ID: primitive.NewObjectID(), FirstName: req.Firstname, LastName: req.Lastname, PhoneNumber: *req.PhoneNumber, Email: req.Email, Password: req.Password}

	result, err := db.CreateUser(user)
	if err != nil {
		log.Fatalf("cannot create user :%v", err)
	}
	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Fatalf("failed to cast ObjectID")
	}

	token, err := access.CreateToken(map[string]string{"id": id.Hex(), "email": req.Email})
	if err != nil {
		log.Fatalf("cannot create token:%v", err)
	}
	log.Println("acess token", token)
	return &pb.AuthUser{
		Id:          id.Hex(),
		Firstname:   req.Firstname,
		Lastname:    req.Lastname,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		AccessToken: token,
	}, nil
}
