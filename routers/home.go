package routers

import (
	"net/http"
	"fmt"
)

func NewRouter() http.Handler {
	main := http.NewServeMux()
	main.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)

		if r.URL.Path == "/createUser" {
			res := CreateUsers()
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(res))
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