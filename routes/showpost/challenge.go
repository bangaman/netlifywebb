package showpost

import (
	"net/http"
	"mungo"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
	"templates"
	"html/template"
    "fmt"
    "strings"
)

type ChallengeStruct struct{
	ChallengePost template.HTML
}

func Challenge(w http.ResponseWriter, r *http.Request){
	let := ChallengeStruct{}
    session, _ := mungo.Store.Get(r, "session")

    if len(r.URL.Path[len("/challenge/"):]) > 0 {
        getcoll := mungo.Connect("app").Collection("post")
        getquestion := mungo.Connect("app").Collection("questions")
        vax := strings.Split(r.URL.Path[len("/challenge/"):], "/")

        if len(session.Values["username"].(string)) > 1{
			theid, _:= primitive.ObjectIDFromHex(vax[0])
            result := mungo.FindOne(getcoll, bson.M{"_id":theid})
            question := mungo.FindOne(getquestion, bson.M{"_id":theid})

			if len(result) > 0 {
				let.ChallengePost = template.HTML(strings.Join(GetAllPosts(vax[0], session.Values["usersid"].(string)), ""))
                w.Header().Set("Content-Type", "text/html")
               fmt.Fprintf(w, templates.HTML(let, "https://maajichallenger.netlify.app/templates/showpost/challenge"))
			}else if len(question) > 0 {
                let.ChallengePost = template.HTML(strings.Join(GetAllQuestions(vax[0], session.Values["usersid"].(string)), ""))
                 w.Header().Set("Content-Type", "text/html")
                fmt.Fprintf(w, templates.HTML(let, "https://maajichallenger.netlify.app/templates/showpost/challenge"))
            }else{
				fmt.Fprintf(w, "No challenge to show")
			}
               
        }else{
            http.Redirect(w, r, "/login", http.StatusSeeOther)
        }

    }else{
        http.Redirect(w, r, "/", http.StatusSeeOther)
    }
}
