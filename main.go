package main

import (
	"context"
	"fmt"
	"net/url"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	// Data source
	driver := "mongodb"
	host := "127.0.0.1"
	port := "27017"
	username := "kecci"
	password := "p@sword%20"

	// Unescape username
	username, err := url.QueryUnescape(username)
	if err != nil {
		fmt.Println(err)
	}

	// Unescape password
	password, err = url.QueryUnescape(password)
	if err != nil {
		fmt.Println(err)
	}

	// Init URL
	URL := url.URL{
		Scheme: driver,
		Host:   fmt.Sprintf("%s:%s", host, port),
	}

	// Set username & password
	if username != "" || password != "" {
		URL.User = url.UserPassword(username, password)
	}

	// Set URI
	clientOptions := options.Client().ApplyURI(URL.String())
	clientOptions.SetDirect(true)

	// Connect
	ctx := context.Background()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Printf("error when mongo.Connect(ctx, clientOptions), %v, %v", clientOptions, err)
	}

	// Ping Connextion
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Success
	fmt.Println("ping success")

	// Close context
	<-ctx.Done()
}
