package register

import (
	"net/http"
	"templates"
	"fmt"
	"html/template"
)


func Register(w http.ResponseWriter, r *http.Request){
	let := RegisterStruct{}
	let.RegisterProcess(r)
	
	if let.Something == true{
		fmt.Fprintf(w, "%s", template.HTML("<h3>Something happened while registering, <a href='/bbm/register'>Go back</a></h3>"))
	}else{
		if let.Succes == true{
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}else{
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprintf(w, templates.HTML(let, "https://maajichallenger.netlify.app/templates/register/register"))
		}
		
	}
}