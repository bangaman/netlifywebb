package logout


import (
	"net/http"
	"mungo"
)

func Logout(w http.ResponseWriter, r *http.Request){
	session, _ := mungo.Store.Get(r, "session")
	session.Values["username"] = ""
	session.Values["email"] = ""
	session.Values["usersid"] = ""
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
























