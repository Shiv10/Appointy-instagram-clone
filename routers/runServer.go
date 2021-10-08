package routers

import (
	"fmt"
	"net/http"
	"go.mongodb.org/mongo-driver/mongo"
	"context"
)

func RunHttp(listenAddr string, ctx context.Context, client *mongo.Client) error {
	s := http.Server{
	  Addr:    listenAddr,
	  Handler: NewRouter(client, ctx),
	}
	fmt.Printf("Starting HTTP listener at %s\n", listenAddr)
	return s.ListenAndServe()
  }