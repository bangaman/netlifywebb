package mungo

import (
	"context"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var (
	serv = "mongodb+srv://test:22438802@cluster0.1vm9c.mongodb.net/test?retryWrites=true&w=majority"
	loc = "mongodb://localhost:27017"
)


func Connect(db string)*mongo.Database{
	client, err := mongo.NewClient(options.Client().ApplyURI(serv))
   
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client.Database(db)
}

