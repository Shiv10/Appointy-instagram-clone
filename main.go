package main

import (
	"Appointy-instagram-clone/routers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"context"
	"time"
	"net"
	"flag"
)


func main(){

	//Setup Mongo connection
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://appointy:1234567890@cluster0.1yuhs.mongodb.net/AppointyDB?retryWrites=true&w=majority"))
	if err!= nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1000000*time.Second)
	err = client.Connect(ctx)
	if err!= nil {
		log.Fatal(err)
	}
	defer cancel()
	defer client.Disconnect(ctx)

	//HTTP Sever setup
	var (
		host = flag.String("host", "", "host http address to listen on")
		port = flag.String("port", "3000", "port number for http listener")
	)

	flag.Parse()
	addr := net.JoinHostPort(*host, *port)
	if err := routers.RunHttp(addr, ctx, client); err != nil {
		log.Fatal(err)
	}
}