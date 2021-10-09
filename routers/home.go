package routers

import (
	"net/http"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"context"
	"strings"
	"encoding/json"
)

type Error struct {
	Err string
}

var ctxx context.Context
var cli *mongo.Client

func NewRouter(client *mongo.Client, ctx context.Context) http.Handler {
	ctxx = ctx
	cli = client
	main := http.NewServeMux()
	main.HandleFunc("/", HandleRequest)
	return main
  }

func HandleRequest(w http.ResponseWriter, r *http.Request) {

	ctx := ctxx
	client := cli

	if r.URL.Path == "/user" && r.Method == http.MethodPost {
		res := CreateUsers(client, ctx, w, r)
		fmt.Println(res)
		return
	}

	if strings.Contains(r.URL.Path,"/users") && !strings.Contains(r.URL.Path,"/posts") && r.Method == http.MethodGet {
		id_arr := strings.Split(r.URL.Path, "/")
		id := id_arr[len(id_arr)-1]
		
		if id=="users" {
			var e Error
			e.Err = "User ID missing"
			responseJson, err := json.Marshal(e)
			if err != nil{
				http.Error(w, "Internal Error", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type","application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(responseJson)
			return
		}
		res := GetUser(client, ctx, w, r, id)
		fmt.Println(res)
		return
	}

	if r.URL.Path == "/posts" && r.Method == http.MethodPost {
		res := CreatePost(client, ctx, w, r)
		fmt.Println(res)
		return
	}

	if strings.Contains(r.URL.Path,"/posts") && !strings.Contains(r.URL.Path,"/users") && r.Method == http.MethodGet {
		postID_arr := strings.Split(r.URL.Path, "/")
		postID := postID_arr[len(postID_arr)-1]
		
		if postID=="posts" {
			var e Error
			e.Err = "Post ID missing"
			responseJson, err := json.Marshal(e)
			if err != nil{
				http.Error(w, "Internal Error", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type","application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(responseJson)
			return
		}
		res := GetPost(client, ctx, w, r, postID)
		fmt.Println(res)
		return
	}

	if strings.Contains(r.URL.Path,"/posts/users") && r.Method == http.MethodGet {
		id_arr := strings.Split(r.URL.Path, "/")
		id := id_arr[len(id_arr)-1]
		
		if id=="users" {
			var e Error
			e.Err = "User ID missing"
			responseJson, err := json.Marshal(e)
			if err != nil{
				http.Error(w, "Internal Error", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type","application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(responseJson)
			return
		}
		
		if r.URL.Query().Get("page")!=""{
			tempID := strings.Split(id, "?")
			id = tempID[0]
		}

		res := GetPostByUser(client, ctx, w, r, id)
		fmt.Println(res)
		return
	}

	if r.URL.Path != "/" || r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}