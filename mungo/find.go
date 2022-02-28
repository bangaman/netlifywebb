package mungo


import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAll(coll *mongo.Collection, d bson.M) []bson.M {
	// bson.M{}
	getall, err := coll.Find(context.TODO(), d)

	if err != nil {
		log.Fatal(err)
	}

	var collector []bson.M

	if err = getall.All(context.TODO(), &collector); err != nil {
		log.Fatal(err)
	}

	return collector
}

func FindOne(coll *mongo.Collection, d bson.M) bson.M {
	// opts := options.FindOne().SetSort(bson.D{{"age", 1}})
	var result bson.M
	err := coll.FindOne(
		context.TODO(), d,
		//opts => can also be added
	).Decode(&result)

	if err != nil {
		//if no document matched
		// if err == mongo.ErrNoDocuments {
		// 	return 
		// }
		return result
		// log.Fatal(err)
		//i prefer to check length tha errors
	}

	return result
	//to find result length wrap => len(result)
}


func FindAndDelete(coll *mongo.Collection, id string) bson.M {
	// opts := options.FindOneAndDelete().SetProjection( setpr )
	//opts can be added
	var deleteDocs bson.M
	ids, _ := primitive.ObjectIDFromHex(id)
	err := coll.FindOneAndDelete(context.TODO(), bson.D{{ "_id", ids }} ).Decode(&deleteDocs)

	if err != nil {
		log.Fatal(err)
	}

	return deleteDocs
}



func FindAndUpdate(coll *mongo.Collection, d bson.M, id string) bson.M {
	opts := options.FindOneAndUpdate().SetUpsert(false)
	ids, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{ "_id", ids }}
	update := bson.D{{"$set", d }}
	var updateDocs bson.M

	err := coll.FindOneAndUpdate(context.TODO(),filter, update, opts).Decode(&updateDocs)

	if err != nil {
		log.Fatal(err)
	}

	return updateDocs
}

func QueryJoin(coll *mongo.Collection) []bson.M {
	lookupStorage := bson.D{{"$lookup", bson.D{{ "from", "register" }, {"localField", "usersid"}, {"foreignField", "_id"}, {"as", "theid"} }}}
	unwindStage := bson.D{{"$unwind", bson.D{{ "path", "$usersid" }, {"preserveNullAndEmptyArrays", false} }}}

	showall, err := coll.Aggregate(context.TODO(), mongo.Pipeline{lookupStorage, unwindStage})

	if err != nil {
		panic(err)
	}

	var showLoaded []bson.M

	if err = showall.All(context.TODO(), &showLoaded); err != nil {
		panic(err)
	}

	return showLoaded
}


func QueryJoinPublish(coll *mongo.Collection) []bson.M {
	lookupStorage := bson.D{{"$lookup", bson.D{{ "from", "register" }, {"localField", "usersid"}, {"foreignField", "_id"}, {"as", "theid"} }}}
	unwindStage := bson.D{{"$unwind", bson.D{{ "path", "$usersid" }, {"preserveNullAndEmptyArrays", true} }}}

	showall, err := coll.Aggregate(context.TODO(), mongo.Pipeline{lookupStorage, unwindStage})

	if err != nil {
		panic(err)
	}

	var showLoaded []bson.M

	if err = showall.All(context.TODO(), &showLoaded); err != nil {
		panic(err)
	}

	return showLoaded
}



func QueryJoinSpace(coll *mongo.Collection) []bson.M {
	lookupStorage := bson.D{{"$lookup", bson.D{{ "from", "space" }, {"localField", "subscription"}, {"foreignField", "_id"}, {"as", "theid"} }}}
	unwindStage := bson.D{{"$unwind", bson.D{{ "path", "$subscription" }, {"preserveNullAndEmptyArrays", true} }}}

	showall, err := coll.Aggregate(context.TODO(), mongo.Pipeline{lookupStorage, unwindStage})

	if err != nil {
		panic(err)
	}

	var showLoaded []bson.M

	if err = showall.All(context.TODO(), &showLoaded); err != nil {
		panic(err)
	}

	return showLoaded
}