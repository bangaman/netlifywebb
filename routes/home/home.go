package home


import (
   "fmt"
   "net/http"
   "templates"
   "mungo"
   "showpost"
   "html/template"
   "strings"
)

type HomeStruct struct{
      Username string
      Profile string
      HomePost template.HTML
      HomePublish template.HTML
}

func Home(w http.ResponseWriter, r *http.Request){
      let := HomeStruct{}
      getpost := showpost.ShowPostStruct{}
      session, _ := mungo.Store.Get(r, "session")

      if len(session.Values["username"].(string)) > 1{
            let.Username = session.Values["username"].(string)
            let.Profile = let.Username[0:1]

            getpost.HomeChallenge(session.Values["usersid"].(string))
            let.HomePost = template.HTML(strings.Join(getpost.HomePosts, ""))
            
            w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, templates.HTML(let, "https://maajichallenger.netlify.app/templates/home/home"))
      }else{
            http.Redirect(w, r, "/login", http.StatusSeeOther)
      }
}



func HomeQuestion(w http.ResponseWriter, r *http.Request){
      let := HomeStruct{}
      getpost := showpost.ShowPostStruct{}
      session, _ := mungo.Store.Get(r, "session")

      if len(session.Values["username"].(string)) > 1{
            let.Username = session.Values["username"].(string)
            let.Profile = let.Username[0:1]
            getpost.HomeQuestion(session.Values["usersid"].(string))
            let.HomePost = template.HTML(strings.Join(getpost.HomePosts, ""))
            templates.Render(w, "site/templates/home/question.html", let)
            // w.Header().Set("Content-Type", "text/html")
		// fmt.Fprintf(w, templates.HTML(let, "https://maajichallenger.netlify.app/templates/register/login"))
      }else{
            http.Redirect(w, r, "/login", http.StatusSeeOther)
      }
}

