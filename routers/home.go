package routers

import (
	"net/http"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"context"
)

func NewRouter(client *mongo.Client, ctx context.Context) http.Handler {
	main := http.NewServeMux()
	main.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)

		if r.URL.Path == "/user" && r.Method == http.MethodPost {
			res := CreateUsers(client, ctx, w, r)
			fmt.Println(res)
			return
		}
		if r.URL.Path == "/getUser" {
			res := GetUser()
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(res))
			return
		}

		if r.URL.Path == "/createPost" {
			res := CreatePost()
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(res))
			return
		}

		if r.URL.Path == "/getPost" {
			res := GetPost()
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(res))
			return
		}

		if r.URL.Path == "/getPostByUser" {
			res := GetPostByUser()
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(res))
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