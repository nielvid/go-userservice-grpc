package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/nielvid/go-userservice-grpc/models"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	  // Check the connection.
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }else{
      fmt.Println("Connected to mongoDB!!!")
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

func (db *DB) FindUsers() ([]models.User, error) {
	col := db.client.Database(DBName).Collection(collection)
	// a stream of documents on which we can iterate or we can get all the docs by function cursor.All() into a slice type.
	cursor, err := col.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var users []models.User
	if err = cursor.All(context.TODO(), &users); err != nil {
		log.Fatal(err)
	}
	return users, nil
}

func (db *DB) FindUser(id string) (models.User, error) {

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	filter := bson.D{{"_id", objID}}
	var user models.User
	col := db.client.Database(DBName).Collection(collection)
	err = col.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	return user, nil
}


func (db *DB) UpdateUser(id string, data interface{}) (string, error) {

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	filter := bson.D{{"_id", objID}} 
	col := db.client.Database(DBName).Collection(collection)
	result, err := col.UpdateOne(context.TODO(), filter, bson.M{
		"$set": bson.M{
            "body":       "Some updated text",
            "updated_at": time.Now(),
            },
	})
	if err != nil {
		log.Fatal(err)
	}
	return result.UpsertedID.(primitive.ObjectID).Hex(), nil
}

func (db *DB)DeleteUser(id string) {

    // Delete one document.
    objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
		col := db.client.Database(DBName).Collection(collection)
    result, err := col.DeleteOne(context.TODO(), bson.M{"_id": objID})
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(result.DeletedCount) // output: 1
}

