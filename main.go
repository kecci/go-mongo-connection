package main

import (
	"context"
	"fmt"
	"net/url"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var ctx = context.Background()

func main() {
	driver := "mongodb"
	host := "127.0.0.1"
	port := "27017"
	username, err := url.QueryUnescape("kecci")
	if err != nil {
		username = "kecci"
	}
	password, err := url.QueryUnescape("p@sword%20")
	if err != nil {
		password = "p@sword%20"
	}

	URL := url.URL{
		Scheme: driver,
		Host:   fmt.Sprintf("%s:%s", host, port),
	}

	if username != "" || password != "" {
		URL.User = url.UserPassword(username, password)
	}

	fmt.Printf("URI: %s \n", URL.String())
	clientOptions := options.Client().ApplyURI(URL.String())
	clientOptions.SetDirect(true)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Printf("error when mongo.Connect(ctx, clientOptions), %v, %v", clientOptions, err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("ping success")
}
