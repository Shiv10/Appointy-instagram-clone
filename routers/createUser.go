package routers

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"net/http"
	"encoding/json"
	"crypto/sha1"
)

type User struct {
	Username string
	Password string
	Email string
}

type UsersFromDB struct {
	ID primitive.ObjectID
	Name string
	Email string
	password []uint8
}

type UserResponse struct {
	UserID string
}

func CreateUsers(client *mongo.Client, ctx context.Context, w http.ResponseWriter, r *http.Request) string{
	
	u := User{}
	err := json.NewDecoder(r.Body).Decode(&u)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return "Error"
    }

	appointyDB := client.Database("AppointyDB")
	usersCollection := appointyDB.Collection("Users")

	//Test if user already exists
	filter := bson.D{{Key: "Email", Value: u.Email}}

	result := UsersFromDB{}
	err = usersCollection.FindOne(ctx, filter).Decode(&result)
	if err!=  mongo.ErrNoDocuments {
		var e Error
		e.Err = "User already exists"
		responseJson, err := json.Marshal(e)
		if err != nil{
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return "Error"
		}
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJson)
		return "User exists"
	}

	if err!=nil && err!= mongo.ErrNoDocuments {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
        return "Error"
	}

	//hash the password
	h := sha1.New()
	h.Write([]byte(u.Password))
	bs := h.Sum(nil)

	//insert user in DB
	newUser, err := usersCollection.InsertOne(ctx, bson.D{
		{Key: "Name", Value: u.Username},
		{Key: "Password", Value: bs},
		{Key: "Email", Value: u.Email},
		{Key: "ID", Value: "null"},
	})

	

	if err!=nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
        return "Error"
	}

	stringID := newUser.InsertedID.(primitive.ObjectID).Hex()

	updatedUser, err := usersCollection.UpdateOne(ctx, bson.M{"Email": u.Email}, bson.D{
		{Key: "$set", Value: bson.D{{Key: "ID", Value: stringID}}},
	})

	fmt.Println(updatedUser)
	var userResponse UserResponse
	userResponse.UserID = stringID
	userJson, err := json.Marshal(userResponse)
	if err != nil{
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return "No document found"
	}
    w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK) 
	w.Write([]byte(userJson))

	
	return "User created."
}