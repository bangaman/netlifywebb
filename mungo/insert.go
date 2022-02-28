package mungo


import (
	"context"
	"log"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)


func InsertMany(coll *mongo.Collection,  d []interface{}){

	opts := options.InsertMany().SetOrdered(false)
	res, err := coll.InsertMany(context.TODO(), d, opts)

	if err != nil {
		log.Fatal("InsertMany error ", err)
	}

	fmt.Println("inserted documents with the IDs %v\n", res.InsertedIDs)
}

func InsertOne(coll *mongo.Collection, d bson.M){
	res, err := coll.InsertOne(context.TODO(), d)

	if err != nil {
		log.Fatal("InsertMany error ", err)
	}

	fmt.Println("inserted documents with the IDs %v\n", res.InsertedID)	
}