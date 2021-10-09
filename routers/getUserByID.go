package routers

import (
	"context"
	"net/http"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type ResponseUser struct {
	Email string
	Name string
	ID string
}

func GetUser(client *mongo.Client, ctx context.Context, w http.ResponseWriter, r *http.Request, id string) string{


	appointyDB := client.Database("AppointyDB")
	usersCollection := appointyDB.Collection("Users")
	filter := bson.D{{Key: "ID", Value: id}}

	result := UsersFromDB{}
	err := usersCollection.FindOne(ctx, filter).Decode(&result)
	if err!=  mongo.ErrNoDocuments {
		w.Header().Set("Content-Type", "application/json")
		resp := ResponseUser{}
		resp.Email = result.Email
		resp.Name = result.Name
		resp.ID = id
		userJson, err := json.Marshal(resp)
		if err != nil{
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return "error"
		}
		w.WriteHeader(http.StatusOK)
		w.Write(userJson)
        return "User Sent"
	}

	if err!=nil && err!= mongo.ErrNoDocuments {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
        return "Error"
	}

	var e Error
	e.Err = "No user found for user ID"
	responseJson, err := json.Marshal(e)
	if err != nil{
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return "Error"
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
	return "No user found"
}