package mungo


import (

	"context"
	"log"
	// "time"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func UpdateOne(coll *mongo.Collection, d bson.M, id string){
	ids, _ := primitive.ObjectIDFromHex(id)
	opts := options.Update().SetUpsert(false)
	filter := bson.D{{"_id", ids}}
	update := bson.D{{"$set", d }}

	result, err := coll.UpdateOne(context.TODO(), filter, update, opts)

	if err != nil {
		log.Fatal(err)
	}

	if result.MatchedCount != 0 {
		fmt.Println("Matched and replaced an existing document")
		return
	}

	if result.UpsertedCount != 0{
		fmt.Printf("inserted a new document with ID %v\n", result.UpsertedID)
	}else{
		fmt.Printf("nothing Updated\n")
	}
}