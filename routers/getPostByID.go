package routers

import (
	"context"
	"net/http"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type PostsFromDB struct {
	_ID primitive.ObjectID
	Caption string
	ImageURL string
	UserID string
	ID string
}

type PostResponse struct {
	Caption string
	ImageURL string
	UserID string
	ID string
}

func GetPost(client *mongo.Client, ctx context.Context, w http.ResponseWriter, r *http.Request, postID string) string{

	appointyDB := client.Database("AppointyDB")
	postsCollection := appointyDB.Collection("Posts")
	filter := bson.D{{Key: "ID", Value: postID}}

	var posts []PostResponse
	cur, err := postsCollection.Find(ctx, filter)

	if err==mongo.ErrNoDocuments {
		http.Error(w, "No post found", http.StatusOK)
    	return "Not found"
	}

	if err!=nil && err!= mongo.ErrNoDocuments {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
        return "Error"
	}

	for cur.Next(ctx){
		var elem PostsFromDB
		var resp PostResponse
		err := cur.Decode(&elem)
		if err!=nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return "Error"
		}
		resp.Caption = elem.Caption
		resp.ImageURL = elem.ImageURL
		resp.ID = elem.ID
		resp.UserID = elem.UserID
		posts = append(posts, resp)
	}
	
	if len(posts)==0{
		http.Error(w, "No posts for given id", http.StatusInternalServerError)
		return "No posts"
	}
	postsJson, err := json.Marshal(posts)
	if err != nil{
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return "error"
	}
	w.WriteHeader(http.StatusOK)
	w.Write(postsJson)
	return "post sent."
}