package routers

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"context"
	"net/http"
	"encoding/json"
	"time"
)

type Post struct {
	Caption string
	ImageURL string
	UserID string
}

type PostResp struct {
	PostID string
}

func CreatePost(client *mongo.Client, ctx context.Context, w http.ResponseWriter, r *http.Request) string{
	post := Post{}
	err := json.NewDecoder(r.Body).Decode(&post)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return "Error"
    }

	appointyDB := client.Database("AppointyDB")
	postsCollection := appointyDB.Collection("Posts")
	usersCollection := appointyDB.Collection("Users")


	//test if user with given ID exists or not
	filter := bson.D{{Key: "ID", Value: post.UserID}}

	result := UsersFromDB{}
	err = usersCollection.FindOne(ctx, filter).Decode(&result)

	if err==mongo.ErrNoDocuments {
		var e Error
		e.Err = "No user with ID exists"
		responseJson, err := json.Marshal(e)
		if err != nil{
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return "Error"
		}
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJson)
		return "User does not exist"
	}
	
	if err!=  mongo.ErrNoDocuments {
		newPost, err := postsCollection.InsertOne(ctx, bson.D{
			{Key: "Caption", Value: post.Caption},
			{Key: "UserID", Value: post.UserID},
			{Key: "ID", Value: "null"},
			{Key: "ImageUrl", Value: post.ImageURL},
			{Key: "Timestamp", Value: time.Now()},

		})
		if err!=nil{
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return "Error"
		}
		postID := newPost.InsertedID.(primitive.ObjectID).Hex()

		updatedPost, err := postsCollection.UpdateOne(ctx, bson.M{"_id": newPost.InsertedID}, bson.D{
			{Key: "$set", Value: bson.D{{Key: "ID", Value: postID}}},
		})
		fmt.Println(updatedPost)
		var postResp PostResp
		postResp.PostID = postID
		postJson, err := json.Marshal(postResp)
		if err != nil{
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return "No document found"
		}
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(postJson))
	}

	if err!=nil{
		http.Error(w, "Internal Error", http.StatusInternalServerError)
        return "Error"
	}

	return "Post created."
}