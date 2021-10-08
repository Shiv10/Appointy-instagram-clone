package main

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	// "go.mongodb.org/mongo-driver/bson"
	"log"
	"context"
	"time"
	// "fmt"
	"os"
)

func main(){
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://appointy:1234567890@cluster0.1yuhs.mongodb.net/AppointyDB?retryWrites=true&w=majority"))
	if err!= nil {
		log.Fatal(err)
		os.Exit(1)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err!= nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer cancel()
	defer client.Disconnect(ctx)

	appointyDB := client.Database("AppointyDB")
	usersCollection := appointyDB.Collection("Users")
	postsCollection := appointyDB.Collection("Posts")


}