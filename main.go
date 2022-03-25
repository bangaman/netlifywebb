package main


import (
    "net/http"
	"flag"
	"log"
    "fmt"
    "home"
	"register"
	"os"
	"github.com/carlmjohnson/gateway"
	// "mungo"
	"edit"
	"showpost"
	"logout"
	"post"
)
 



func main() {
	path, _ := os.Getwd()
	fmt.Println(path)
	port := flag.Int("port", -1, "specify a port to use http rather than AWS Lambda")
	flag.Parse()
	listener := gateway.ListenAndServe
	portStr := ""
	if *port != -1 {
		portStr = fmt.Sprintf(":%d", *port)
		listener = http.ListenAndServe
	}

	http.Handle("/site/", http.StripPrefix("/site/", http.FileServer(http.Dir("site"))))
	http.HandleFunc("/", home.Home)
	http.HandleFunc("/q", home.HomeQuestion)
	http.HandleFunc("/login", register.Login)
	http.HandleFunc("/logout", logout.Logout)
	http.HandleFunc("/register", register.Register)
	http.HandleFunc("/challenge/", showpost.Challenge)
	http.HandleFunc("/edit/medal/", edit.EditMedal)
	http.HandleFunc("/edit/publish/", edit.EditPublish)
	http.HandleFunc("/edit/join/", edit.EditJoin)
	http.HandleFunc("/edit/choice/", edit.EditChoice)
	http.HandleFunc("/edit/q/", edit.EditQuestion)
	http.HandleFunc("/make/p/", post.Post)
	http.HandleFunc("/make/q/", post.PostQuestion)
	log.Fatal(listener(portStr, nil))
}

func cacheControlMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=300")
		h.ServeHTTP(w, r)
	})
}
