package showpost

import (
	"mungo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"strings"
	"fmt"
)

func FindJoin(postid, id string ) string {
	getcoll := mungo.Connect("app").Collection("join")
	getter := mungo.FindOne(getcoll, bson.M{"usersid":postid, "id":id})

	if len(getter) > 0 {
		return "Joined"
	}else{
		return "Join"
	}

	return ""

}


func CountJoin(postid string ) string {
	getcoll := mungo.Connect("app").Collection("join")
	getter := mungo.FindAll(getcoll, bson.M{})

	counter := []string{}

	for _, item := range getter {

		if postid == item["usersid"].(string){
			counter = append(counter, item["usersid"].(string))
		}
	}

	return strconv.Itoa(len(counter))
}


func CountMedals(postid, id string ) string {
	getcoll := mungo.Connect("app").Collection("medal")
	getter := mungo.FindAll(getcoll, bson.M{})

	counter := []string{}

	usercheck := []string{}

	result := []string{}

	for _, item := range getter {

		if postid == item["usersid"].(string){
			counter = append(counter, item["usersid"].(string))
		}

		if id == item["id"].(string) && postid == item["usersid"].(string) {
			usercheck = append(usercheck, "liked")
		}
	}

	fmt.Println(len(counter))

	if len(usercheck) > 0 {
		result = append(result, "<i style='pointer-events:none;color:red;'  class='fas fa-medal'></i>")
	}else{
		result = append(result, "<i style='pointer-events:none;'  class='fas fa-medal'></i>")
	}

	if len(counter) > 0 {
		result = append(result, " <b>"+strconv.Itoa(len(counter))+"</b>")
	}else{
		result = append(result, " <b>0</b>")
	}

	return strings.Join(result, " ")
}


func CheckIfUserJoined(postid, id string ) string {
	getcoll := mungo.Connect("app").Collection("join")
	getter := mungo.FindAll(getcoll, bson.M{})

	counter := []string{}

	for _, item := range getter {

		if id == item["id"].(string) && postid == item["usersid"]{
			counter = append(counter, item["usersid"].(string))
		}
	}

	if len(counter) > 0 {
		return "style='display:block;'"
	}else{
		return "style='display:none;'"
	}
}

func CheckUserPublished(postid, id string) string {
	getc := mungo.Connect("app").Collection("publish")
	stuff := mungo.QueryJoinPublish(getc)

	PublishPosts := []string{}

	for _, item := range stuff {

		if postid == item["postid"].(string) {
			PublishPosts = append(PublishPosts, "<div class='publish-section' id='publish-"+postid+"'><div class='child-publish-section'>")
			//open top details div
			PublishPosts = append(PublishPosts, "<div class='publish-top-details'>")
			PublishPosts = append(PublishPosts,  "<div class='publish-section-image'>")
			PublishPosts = append(PublishPosts, "<div id='publish-pic'><button>"+item["theid"].(primitive.A)[0].(primitive.M)["username"].(string)[0:1]+"</button></div>")
			PublishPosts = append(PublishPosts, "<div id='publish-username'><span><a href='/user/"+item["theid"].(primitive.A)[0].(primitive.M)["username"].(string)+"'>"+item["theid"].(primitive.A)[0].(primitive.M)["username"].(string)+"</a></span></div>")
			PublishPosts = append(PublishPosts, "</div>")		
			PublishPosts = append(PublishPosts,  "<div class='publish-section-medal'>")

			if FindJoin(postid, item["usersid"].(primitive.ObjectID).Hex()) == "Join" {
				PublishPosts = append(PublishPosts, "<span><button data-name='"+item["_id"].(primitive.ObjectID).Hex()+"'><div id='medal"+item["_id"].(primitive.ObjectID).Hex()+"'>"+CountMedals(item["_id"].(primitive.ObjectID).Hex(), id)+" <i class='fas fa-window-close'></i></div></button></span>")
				PublishPosts = append(PublishPosts, "<span style='padding-left:5px;'><button>Left Challenge</button></span>")
			}else{
				PublishPosts = append(PublishPosts, "<span><button id='medal-btn' data-name='"+item["_id"].(primitive.ObjectID).Hex()+"'><div id='medal"+item["_id"].(primitive.ObjectID).Hex()+"' style='pointer-events:none;'>"+CountMedals(item["_id"].(primitive.ObjectID).Hex(), id)+"</div></button></span>")
			}
			PublishPosts = append(PublishPosts, "</div>")
	
			//close top details div	
			PublishPosts = append(PublishPosts, "</div>")

			PublishPosts= append(PublishPosts, "<div class='publish-body'><span>"+item["data"].(string)+"</span></div>")
		
			PublishPosts = append(PublishPosts, "</div></div>")
		}
		

	}

	return strings.Join(PublishPosts, "")
}


func FindPublished(postid, id string ) string {
	getcoll := mungo.Connect("app").Collection("publish")
	theid, _:= primitive.ObjectIDFromHex(id)
	getter := mungo.FindOne(getcoll, bson.M{"postid":postid, "usersid":theid})

	if len(getter) > 0 {
		return "published"
	}
	return ""
}



func CheckUserAnsweredQuestion(postid, option, id string) string {
	getc := mungo.Connect("app").Collection("support")
	stuff := mungo.QueryJoinPublish(getc)

	PublishPosts := []string{}

	for _, item := range stuff {

		if postid == item["postid"].(string) {
			PublishPosts = append(PublishPosts, "<div class='publish-section' id='publish-"+postid+"'><div class='child-publish-section'>")
			//open top details div
			PublishPosts = append(PublishPosts, "<div class='publish-top-details'>")
			PublishPosts = append(PublishPosts,  "<div class='publish-section-image'>")
			PublishPosts = append(PublishPosts, "<div id='publish-pic'><button>"+item["theid"].(primitive.A)[0].(primitive.M)["username"].(string)[0:1]+"</button></div>")
			PublishPosts = append(PublishPosts, "<div id='publish-username'><span><a href='/user/"+item["theid"].(primitive.A)[0].(primitive.M)["username"].(string)+"'>"+item["theid"].(primitive.A)[0].(primitive.M)["username"].(string)+"</a></span></div>")
			PublishPosts = append(PublishPosts, "</div>")	

			if option == item["option"].(string){
				PublishPosts = append(PublishPosts, "<div class='option-selected'><span>Selceted <b style='text-transform:uppercase;'>"+item["option"].(string)+"</b>  <i style='color:cadetblue;'>CORRECT</i></span></div>")
			}else{
				PublishPosts = append(PublishPosts, "<div class='option-selected'><span>Selceted <b style='text-transform:uppercase;'>"+item["option"].(string)+"</b>  <i style='color:red;text-transform:uppercase;'>wrong the answer is </i><b style='text-transform:uppercase;'>"+option+"</b></span></div>")
			}
			
			//close top details
			PublishPosts = append(PublishPosts, "</div>")

			
			PublishPosts= append(PublishPosts, "<div class='publish-body'><span>"+item["support"].(string)+"</span></div>")
		
			PublishPosts = append(PublishPosts, "</div></div>")
		}
		

	}

	return strings.Join(PublishPosts, "")
}

func CheckUserAnsweredQuestionOrNot(postid, id string) string {
	getcoll := mungo.Connect("app").Collection("support")
	theid, _:= primitive.ObjectIDFromHex(id)
	getter := mungo.FindOne(getcoll, bson.M{"postid":postid, "usersid":theid})

	if len(getter) > 0 {
		return "true"
	}else{
		return "false"
	}

	return ""
}


func FindChoice(postid, id string ) string {
	getcoll := mungo.Connect("app").Collection("choice")
	getter := mungo.FindOne(getcoll, bson.M{"postid":postid, "usersid":id})

	if len(getter) > 0 {

		if getter["usersid"].(string) == id {
			return getter["choice"].(string)+"T"
		}else{
			return getter["choice"].(string)+"F"
		}

	}else {
		//to manipulate future error
		//if user has not answered question 
		//it still returns something but not what we need
		return "kkkk"
	}
	
	return ""
}


// QueryJoinPublish