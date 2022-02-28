package edit


import (
	"net/http"
	"strings"
	"fmt"
	"mungo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (l *EditStruct) JoinVerify(id string, r *http.Request){
	if strings.ToLower(r.Method) == "post"{
		fmt.Println(r.FormValue("join"))
		getcoll := mungo.Connect("app").Collection("join")
		checkname := len(mungo.FindOne(getcoll, bson.M{"usersid":r.FormValue("join"), "id":id}))
		l.Postid = r.FormValue("join")
		if checkname == 0 {
			mungo.InsertOne(getcoll, bson.M{"usersid":r.FormValue("join"), "id":id})
			
		}else{
			mungo.DeleteOne(getcoll, bson.M{"usersid":r.FormValue("join"), "id":id})
		}
		
	}
}


func (l *EditStruct) MedalVerify(id string, r *http.Request){
	if strings.ToLower(r.Method) == "post"{
		getcoll := mungo.Connect("app").Collection("medal")
		checkname := len(mungo.FindOne(getcoll, bson.M{"usersid":r.FormValue("postid"), "id":id}))

		l.Postid = r.FormValue("postid")
		if checkname == 0 {
			mungo.InsertOne(getcoll, bson.M{"usersid":r.FormValue("postid"), "id":id})
			l.JoinInserted = true
		}else{
			mungo.DeleteOne(getcoll, bson.M{"usersid":r.FormValue("postid"), "id":id})
			l.JoinDeleted = true
		}
		
	}
}


func (l *EditStruct) PublishVerify(id string, r *http.Request){
	if strings.ToLower(r.Method) == "post"{
		getcoll := mungo.Connect("app").Collection("publish")
		checkname := len(mungo.FindOne(getcoll, bson.M{"postid":r.FormValue("postid"), "usersid":id}))
		theid, _:= primitive.ObjectIDFromHex(id)
		l.Postid = r.FormValue("postid")
		if checkname == 0 {
			mungo.InsertOne(getcoll, bson.M{"postid":r.FormValue("postid"), "usersid":theid, "data":r.FormValue("data")})
			l.Published = true
			
		}
	}
}


func (l *EditStruct) QuestionVerify(id string, r *http.Request){
	if strings.ToLower(r.Method) == "post"{
		getcoll := mungo.Connect("app").Collection("support")
		checkname := len(mungo.FindOne(getcoll, bson.M{"postid":r.FormValue("postid"), "usersid":id}))
		theid, _:= primitive.ObjectIDFromHex(id)
		l.Postid = r.FormValue("postid")
		if checkname == 0 {
			mungo.InsertOne(getcoll, bson.M{"postid":r.FormValue("postid"), "usersid":theid, "support":r.FormValue("support"), "option":r.FormValue("option")})
			l.Published = true
			
		}
	}
}


func (l *EditStruct) ChoiceVerify(id string, r *http.Request){
	if strings.ToLower(r.Method) == "post"{
		getcoll := mungo.Connect("app").Collection("choice")
		checkname := len(mungo.FindOne(getcoll, bson.M{"postid":r.FormValue("postid"), "usersid":id}))
		l.Postid = r.FormValue("postid")
		if checkname == 0 {
			mungo.InsertOne(getcoll, bson.M{"postid":r.FormValue("postid"), "usersid":id, "choice":r.FormValue("choice")})
			l.Published = true
			
		}
	}
}