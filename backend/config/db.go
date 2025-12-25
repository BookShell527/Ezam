package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var Client *mongo.Client
var StudentColl *mongo.Collection
var ExamColl *mongo.Collection

func InitializeDB() {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatalf("Please Provide MongoDB URI!!!")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1))

	var err error
	Client, err = mongo.Connect(opts)

	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}

	err = Client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Can't Ping the DB: %s", err)
	}
	StudentColl = Client.Database("ezam").Collection("student")
	ExamColl = Client.Database("ezam").Collection("exam")

	fmt.Println("Connected to MongoDB!")
}
