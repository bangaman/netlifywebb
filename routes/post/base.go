package post

import (
	"net/http"
	"strings"
	"mungo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "fmt"
	"time"
	"strconv"
)

func (l *PostStruct) ProcessPost(id string, r *http.Request){
	if strings.ToLower(r.Method) == "post"{
		getcoll := mungo.Connect("app").Collection("post")
		category := r.FormValue("category")
		thetype := r.FormValue("type")
		thetime := r.FormValue("time")
		tips :=  strings.TrimSpace(r.FormValue("tips"))

		var thebadge string

		if len(r.FormValue("thebatch")) > 0 {
			thebadge = r.FormValue("thebatch")
		}else{
			thebadge = ""
		}
		tag :=  strings.TrimSpace(r.FormValue("tag"))

		theid, _:= primitive.ObjectIDFromHex(id)

		if len(tips) < 1{
			l.Err = true
		}else{
			t := time.Now()
			dateval :=  t.Month().String()[:3]+" "+strconv.Itoa(t.Day())+", "+strconv.Itoa(t.Year())

			mungo.InsertOne(getcoll, bson.M{"question":"false", "usersid":theid, "category":category, "thetype":thetype, "thetime":thetime, "tips":tips, "thebadge":thebadge, "tag":tag, "time":dateval})
		}	
	}
}



func (l *PostStruct) ProcessPostQuestion(id string, r *http.Request){
	if strings.ToLower(r.Method) == "post"{
		getcoll := mungo.Connect("app").Collection("questions")
		question := r.FormValue("q")
		a := r.FormValue("a")
		b := r.FormValue("b")
		c :=  r.FormValue("c")
		d  :=  r.FormValue("d")
		check := r.FormValue("check")
		subject := r.FormValue("subject")
		answer := r.FormValue("default")

		theid, _:= primitive.ObjectIDFromHex(id)

		if len(question) < 1{
			l.Err = true
		}else{
			t := time.Now()
			dateval :=  t.Month().String()[:3]+" "+strconv.Itoa(t.Day())+", "+strconv.Itoa(t.Year())

			mungo.InsertOne(getcoll, bson.M{"question":"true", "usersid":theid, "q":question, "a":a, "b":b, "c":c, "d":d, "check":check, "time":dateval, "subject":subject, "default":answer})
		}

			
	}
}