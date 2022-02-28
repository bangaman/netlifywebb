package register


import(
	"strings"
	"net/http"
	"html/template"
	"fmt"
	"mungo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type RegisterStruct struct{
	Something bool
	Username string
	User_err template.HTML
	Email_err template.HTML
	Email string
	Password template.HTML
	Succes bool
}


func (d *RegisterStruct)  RegisterProcess(r *http.Request){

	if strings.ToLower(r.Method) == "post"{
		geterrors := []string{}
		d.Email = strings.TrimSpace(r.FormValue("email"))
		d.Username = strings.TrimSpace(r.FormValue("username"))
		pass := strings.TrimSpace(r.FormValue("password"))
		
		// defer mungo.Connect().Disconnect(ctx)
		getcoll := mungo.Connect("app").Collection("register")

		spacestarter := mungo.Connect("app").Collection("space")
		
		checkname := len(mungo.FindOne(getcoll, bson.M{"username":strings.ToLower(strings.TrimSpace(r.FormValue("username")))})) 
		checkemail := len(mungo.FindOne(getcoll, bson.M{"email": strings.ToLower(strings.TrimSpace(r.FormValue("email")))})) 

		if len(strings.TrimSpace(r.FormValue("email"))) < 1 || len(strings.TrimSpace(r.FormValue("username"))) < 1 || len(pass) < 1 {
			d.Something = true

		}else{
			if checkname > 0 {
				d.User_err = template.HTML("<div style='padding-top:5px;'><i style='color:red;font-size:14px;'>Username Taken Already</i></div>")
				geterrors = append(geterrors, "username")
			}
	
			if checkemail > 0 {
				d.Email_err = template.HTML("<div style='padding-top:5px;'><i style='color:red;font-size:14px;'>Email Taken Already</i></div>")
				geterrors = append(geterrors, "email")
			}
	
			if len(pass) < 8{
				d.Password = template.HTML("<div style='padding-top:5px;'><i style='color:red;font-size:14px;'>Password Must Be 8 Characters or More</i></div>")
				geterrors = append(geterrors, "password")
			}
	
	
			if len(geterrors) > 0 {
				d.Succes = false
			}else{
				d.Succes = true
				mungo.InsertOne(getcoll, bson.M{"username": strings.ToLower(strings.TrimSpace(r.FormValue("username"))), "email": strings.ToLower(strings.TrimSpace(r.FormValue("email"))), "password": strings.TrimSpace(r.FormValue("password")) })

				checkandcreatespace := mungo.FindOne(getcoll, bson.M{"username": strings.ToLower(strings.TrimSpace(r.FormValue("username")))}) 

				if len(checkandcreatespace) > 0 {
					mungo.InsertOne(spacestarter, bson.M{"name":"starter", "usersid":checkandcreatespace["_id"], "type":"starter"})
				}


			}
			fmt.Println(strings.TrimSpace(r.FormValue("email")), " aha ", strings.TrimSpace(r.FormValue("username")))
		}
		
	}
}


type LoginStruct struct{
	Username string
	Credentials string
	Password string
	Succes bool
	Usersid string
	Email string
}


func (d *LoginStruct)  LoginProcess(r *http.Request){

	if strings.ToLower(r.Method) == "post"{
		geterrors := []string{}
		d.Username = strings.ToLower(strings.TrimSpace(r.FormValue("username")))
		pass := strings.TrimSpace(r.FormValue("password"))
		
		// defer mungo.Connect().Disconnect(ctx)
		getcoll := mungo.Connect("app").Collection("register")
		
		checkcredentials := len(mungo.FindOne(getcoll, bson.M{"username": strings.ToLower(strings.TrimSpace(r.FormValue("username"))), "password":strings.TrimSpace(pass)})) 
		
		if checkcredentials < 1 {
			d.Credentials = "Invalid Credentials"
			geterrors = append(geterrors, "invalid")
		}else{
			d.Usersid = mungo.FindOne(getcoll, bson.M{"username":strings.ToLower(strings.TrimSpace(r.FormValue("username"))), "password":strings.TrimSpace(pass)})["_id"].(primitive.ObjectID).Hex()
			d.Email = mungo.FindOne(getcoll, bson.M{"username":strings.ToLower(strings.TrimSpace(r.FormValue("username"))), "password":strings.TrimSpace(pass)})["email"].(string)
		}
		
		if len(geterrors) > 0 {
			d.Succes = false
		}else{
			d.Succes = true
		}
		// fmt.Println(mungo.FindOne(getcoll, bson.M{"username":strings.TrimSpace(r.FormValue("username")), "password":pass1})["_id"].(primitive.ObjectID).Hex()) 
		fmt.Println(strings.TrimSpace(r.FormValue("email")), " aha ", strings.ToLower(strings.TrimSpace(r.FormValue("username"))))
	}
}