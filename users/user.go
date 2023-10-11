package users

import (
	"context"
	"io"
	"log"
	"time"

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
	db     = database.Connection()
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

// unary
func (s *UserServer) FindUsers(ctx context.Context, req *pb.NoParams) (*pb.Users, error) {
	cursor, err := db.FindUsers()
	if err != nil {
		log.Fatalf("cannot create token:%v", err)
	}
	var users []*pb.User
	for _, value := range cursor {
		res := &pb.User{
			Id:          value.ID.Hex(),
			Firstname:   value.FirstName,
			PhoneNumber: &value.PhoneNumber,
			Email:       value.Email,
		}
		users = append(users, res)
	}

	return &pb.Users{
		Users: users,
	}, nil

}

func (s *UserServer) FindUser(ctx context.Context, req *pb.UserId) (*pb.User, error) {
	user, err := db.FindUser(req.Id)
	if err != nil {
		log.Fatalf("cannot create token:%v", err)
	}

		return  &pb.User{
			Id:         user.ID.Hex(),
			Firstname:   user.FirstName,
			PhoneNumber: &user.PhoneNumber,
			Email:      user.Email,
		}, nil

}

// server streaming i.e server sending strems of messages back to the client
func (s *UserServer) FetchUsers(req *pb.NoParams, stream pb.UserService_FetchUsersServer) error {
	cursor, err := db.FindUsers()
	if err != nil {
		log.Fatalf("cannot create token:%v", err)
	}

	for _, value := range cursor {
		res := &pb.User{
			Id:          value.ID.Hex(),
			Firstname:   value.LastName,
			PhoneNumber: &value.PhoneNumber,
			Email:       value.Email,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}

	return nil
}

// client streaming i.e client sending messages in streams, while server just respond with a message
func (s *UserServer) VerifyUsers(stream pb.UserService_VerifyUsersServer) error {

	var usersId []string

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println(usersId) //to something with the usersId
			return stream.SendAndClose(&pb.VerificationResponse{Message: "request received. Processing started"})
		}
		if err != nil {
			return err
		}
		log.Println(req, "users to verify")
		usersId = append(usersId, req.Id)
	}
}

// bidirectional streaming i.e receive streams and send stream
func (s *UserServer) Chat(stream pb.UserService_ChatServer) error {


	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Println(req, "message received")
		res := &pb.ChatMessage{Message: req.Message}

		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
