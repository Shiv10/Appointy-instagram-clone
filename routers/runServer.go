package routers

import (
	"fmt"
	"net/http"
)

func RunHttp(listenAddr string) error {
	s := http.Server{
	  Addr:    listenAddr,
	  Handler: NewRouter(),
	}
	fmt.Printf("Starting HTTP listener at %s\n", listenAddr)
	return s.ListenAndServe()
  }