package mungo


import (

	"context"
	"log"
	// "time"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)



func DeleteOne(coll *mongo.Collection, d bson.M){
	opts := options.Delete().SetCollation(&options.Collation{
		Locale: "en_US",
		Strength: 1,
		CaseLevel: false,
	})

	res, err := coll.DeleteOne(context.TODO(),d, opts)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("deleted %v document\n", res.DeletedCount)
}
