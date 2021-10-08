package routers

import (
	"net/http"
)

func NewRouter() http.Handler {
	main := http.NewServeMux()
	main.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" || r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
	})
	return main
  }