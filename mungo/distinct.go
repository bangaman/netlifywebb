package mungo

import (

	"context"
	"log"
	"time"
	// "fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

//Distinct executes a distinct command to find the unique values for a specified field in the
//collection


func Distinct(coll *mongo.Collection, d bson.D)[]interface{}{

	//Find all unique values
	//like "age" field is greater than 25

	//specify the MaxTime option to limit the amount of time the operation can 
	//run on the server

	//like d = bson.D{{"age", bson.D{{"$gt", 25}} }}
	//Distinct(coll, bson.D{{"age", bson.D{{"$gt", 25}} }})

	filter :=  d
	opts := options.Distinct().SetMaxTime(2 * time.Second)
	values, err := coll.Distinct(context.TODO(), "name", filter, opts )

	if err != nil {
		log.Fatal(err)
	}


	// fmt.Println(values)
	return values
}

// bson.M{"_id": bson.ObjectIdHex("ghgjhghg") }