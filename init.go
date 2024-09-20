package main

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"log"
	"time"
)

const (
	mongoURI = "mongodb://localhost:27017"
	database = "jktest"
)

var (
	client *mongo.Client
	db     *mongo.Database
)

func init() {
	var err error
	client, err = mongo.Connect(
		options.Client().ApplyURI(mongoURI),
		options.Client().SetMaxPoolSize(100),
		options.Client().SetMaxConnecting(50),
		options.Client().SetMinPoolSize(20),
		options.Client().SetMaxConnIdleTime(120*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("----连接-ok!----")

	db = client.Database(database)
}
