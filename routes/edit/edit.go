package edit 


import (
	"net/http"
	// "templates"
	"mungo"
	"fmt"
	"showpost"
	"strings"
)

type EditStruct struct {
	JoinInserted bool
	JoinDeleted bool
	Postid string
	Published bool
}

func EditJoin(w http.ResponseWriter, r *http.Request){
	session, _ := mungo.Store.Get(r, "session")
	getid := session.Values["usersid"].(string)
	let := EditStruct{}
	if len(session.Values["username"].(string)) > 1{
		let.JoinVerify(getid, r)
		http.Redirect(w, r, "/challenge/"+let.Postid, http.StatusSeeOther)
	}else{
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}


func EditMedal(w http.ResponseWriter, r *http.Request){
	session, _ := mungo.Store.Get(r, "session")
	getid := session.Values["usersid"].(string)
	let := EditStruct{}
	if len(session.Values["username"].(string)) > 1{
		let.MedalVerify(getid, r)

		if let.JoinInserted == true {
			fmt.Fprintf(w, showpost.CountMedals(let.Postid, getid))
		}else{
			fmt.Fprintf(w, showpost.CountMedals(let.Postid, getid))
		}
	}else{
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}


func EditPublish(w http.ResponseWriter, r *http.Request){
	session, _ := mungo.Store.Get(r, "session")
	getid := session.Values["usersid"].(string)
	let := EditStruct{}
	if len(session.Values["username"].(string)) > 1{
		let.PublishVerify(getid, r)
		if let.Published == true {
			fmt.Fprintf(w, "worked")
		}
		
	}else{
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}


func EditQuestion(w http.ResponseWriter, r *http.Request){
	session, _ := mungo.Store.Get(r, "session")
	getid := session.Values["usersid"].(string)
	let := EditStruct{}
	if len(session.Values["username"].(string)) > 1{
		let.QuestionVerify(getid, r)
		if let.Published == true {
			http.Redirect(w, r, "/challenge/"+let.Postid, http.StatusSeeOther)
		}
		
	}else{
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}


func EditChoice(w http.ResponseWriter, r *http.Request){
	session, _ := mungo.Store.Get(r, "session")
	getid := session.Values["usersid"].(string)
	let := EditStruct{}
	if len(session.Values["username"].(string)) > 1{
		let.ChoiceVerify(getid, r)
		if let.Published == true {
			fmt.Fprintf(w, strings.Join(showpost.GetAllQuestions(let.Postid, getid), ""))
		}
		
	}else{
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}