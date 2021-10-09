package routers

import (
	"net/http"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"context"
	"strings"
)

func NewRouter(client *mongo.Client, ctx context.Context) http.Handler {
	main := http.NewServeMux()
	main.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/user" && r.Method == http.MethodPost {
			res := CreateUsers(client, ctx, w, r)
			fmt.Println(res)
			return
		}

		if strings.Contains(r.URL.Path,"/users") && !strings.Contains(r.URL.Path,"/posts") && r.Method == http.MethodGet {
			id_arr := strings.Split(r.URL.Path, "/")
			id := id_arr[len(id_arr)-1]
			
			if id=="users" {
				http.Error(w, "Please Enter user ID", http.StatusBadRequest)
				fmt.Println("No ID")
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
				http.Error(w, "Please Enter post ID", http.StatusBadRequest)
				fmt.Println("No ID")
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
				http.Error(w, "Please Enter user ID", http.StatusBadRequest)
				fmt.Println("No ID")
				return
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
	})
	return main
  }