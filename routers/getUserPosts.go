package routers

import (
	"context"
	"fmt"
	"net/http"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
)

func GetPostByUser(client *mongo.Client, ctx context.Context, w http.ResponseWriter, r *http.Request, id string) string{
	appointyDB := client.Database("AppointyDB")
	postsCollection := appointyDB.Collection("Posts")

	filter := bson.D{{Key: "UserID", Value: id}}

	//Add Pagination
	findOptions := options.Find()
	tempPage := r.URL.Query().Get("page")
	var page int=1
	if tempPage!=""{
		page, _ = strconv.Atoi(tempPage)
	}
	var pageLimit int64 = 5

	findOptions.SetSkip(((int64(page)-1)*pageLimit))
	findOptions.SetLimit(pageLimit)

	var posts []PostResponse
	cur, err := postsCollection.Find(ctx, filter, findOptions)

	if err==mongo.ErrNoDocuments {
		var e Error
		e.Err = "No post found for user ID"
		responseJson, err := json.Marshal(e)
		if err != nil{
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return "Error"
		}
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJson)
		return "No post found"
	}

	if err!=nil && err!= mongo.ErrNoDocuments {
		fmt.Println(err)
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
		resp.Timestamp = elem.Timestamp
		posts = append(posts, resp)
	}
	
	if len(posts)==0{
		var e Error
		e.Err = "No post found"
		responseJson, err := json.Marshal(e)
		if err != nil{
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return "Error"
		}
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJson)
		return "No post found"
	}
	postsJson, err := json.Marshal(posts)
	if err != nil{
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return "Error"
	}

	w.WriteHeader(http.StatusOK)
	w.Write(postsJson)
	return "User posts sent."
}