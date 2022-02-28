package post

import (
	"net/http"
	"templates"
	"mungo"
	"fmt"
)

type PostStruct struct{
	Err bool
}

func Post(w http.ResponseWriter, r *http.Request) {
	session, _ := mungo.Store.Get(r, "session")
	if len(session.Values["username"].(string)) > 1{
		let := PostStruct{}
		let.ProcessPost(session.Values["usersid"].(string),r)
		if let.Err == true{
			fmt.Fprintf(w, "Something happened")
		}else{
			
            w.Header().Set("Content-Type", "text/html")
			fmt.Fprintf(w, templates.HTML(let, "https://maajichallenger.netlify.app/templates/post/post"))
		}
	}else{
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}



func PostQuestion(w http.ResponseWriter, r *http.Request) {
	session, _ := mungo.Store.Get(r, "session")
	if len(session.Values["username"].(string)) > 1{
		let := PostStruct{}
		let.ProcessPostQuestion(session.Values["usersid"].(string),r)
		if let.Err == true{
			fmt.Fprintf(w, "Something happened")
		}else{
			templates.Render(w, "site/templates/post/question.html", nil)
            // w.Header().Set("Content-Type", "text/html")
		// fmt.Fprintf(w, templates.HTML(let, "https://maajichallenger.netlify.app/templates/post/post"))
		}
	}else{
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
