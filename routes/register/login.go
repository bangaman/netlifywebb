package register

import (
	"net/http"
	"templates"
	"mungo"
	"fmt"
)


func Login(w http.ResponseWriter, r *http.Request){
	let := LoginStruct{}
	let.LoginProcess(r)

	if let.Succes == true{
		session, _ := mungo.Store.Get(r, "session")
		session.Values["username"] = let.Username
		session.Values["email"] = let.Email
		session.Values["usersid"] = let.Usersid
		session.Save(r, w)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}else{
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, templates.HTML(let, "https://maajichallenger.netlify.app/templates/register/login"))
	}
}
