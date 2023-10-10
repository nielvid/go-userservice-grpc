package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/nielvid/go-userservice-grpc/models"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DBName     = "userService"
	collection = "users"
)

type DB struct {
	client *mongo.Client
}

func Connection() *DB {
	godotenv.Load()
	uri := os.Getenv("MONGOURI")

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
// 	col := client.Database(DBName).Collection("users")
//create index on email and make it unique
// 	col.Indexes().CreateOne(ctx, mongo.IndexModel{Keys: bson.D{{ "email", 1}},
//  Options: options.Index().SetUnique(true),})
	

	return &DB{client: client}
}

func (db *DB) CreateUser(user *models.User) (*mongo.InsertOneResult, error) {
	col := db.client.Database(DBName).Collection(collection)
	result, err := col.InsertOne(context.TODO(), user)
	return result, err

}
