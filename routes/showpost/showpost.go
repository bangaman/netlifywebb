package showpost

import (
	"mungo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "fmt"
	// "strings"
)

type ShowPostStruct struct{
	GetPosts []string
	UserPosts []string
	HomePosts []string
	PublishPosts []string
}


func (l *ShowPostStruct) GetAllPosts(id string){

	getc := mungo.Connect("app").Collection("post")
	stuff := mungo.QueryJoin(getc)

	

	for _, item := range stuff {

		l.GetPosts = append(l.GetPosts, "<div class='post-section' style='margin-top:15px;'>")
		//open top details div
		l.GetPosts = append(l.GetPosts, "<div class='child-post-section' style='border:1px solid #ccc;'>")
		
		l.GetPosts = append(l.GetPosts, "<div class='post-top-details'>")
		l.GetPosts = append(l.GetPosts,  "<div class='post-section-image'>")
		l.GetPosts = append(l.GetPosts, "<div id='post-pic'><button>"+item["theid"].(primitive.A)[0].(primitive.M)["username"].(string)[0:1]+"</button></div>")
		l.GetPosts = append(l.GetPosts, "<div id='post-username'><span>"+item["theid"].(primitive.A)[0].(primitive.M)["username"].(string)+"</span></div>")
		l.GetPosts = append(l.GetPosts, "</div>")

		// if item["_id"].(primitive.ObjectID).Hex() 
		if item["theid"].(primitive.A)[0].(primitive.M)["_id"].(primitive.ObjectID).Hex() != id {
			l.GetPosts = append(l.GetPosts,  "<div class='post-section-join'>")
			l.GetPosts = append(l.GetPosts, "<button id='join-btn' data-name='"+item["_id"].(primitive.ObjectID).Hex()+"'><span id='join-btn"+item["_id"].(primitive.ObjectID).Hex()+"' style='pointer-events:none;'>"+FindJoin(item["_id"].(primitive.ObjectID).Hex(), id)+"</span></button></div>")
		}

		//close top-details div
		l.GetPosts = append(l.GetPosts, "</div>")

		//show medal if any

		if item["thebadge"] != ""{
			l.GetPosts = append(l.GetPosts, "<div class='post-medal'><div class='child-post-medal'>")
			if item["thebadge"] == "medal"{
				l.GetPosts = append(l.GetPosts, "<i style='color: goldenrod;pointer-events:none;' class='fas fa-medal'></i>")
			}else{
				l.GetPosts = append(l.GetPosts, "<i  style='color: goldenrod;pointer-events:none;'  class='fas fa-trophy'></i>")
			}

			l.GetPosts = append(l.GetPosts, "</div></div>")
		}

		if item["thebadge"] != "" {
			l.GetPosts = append(l.GetPosts, "<div class='post-tagg-medal'><span>#"+item["tag"].(string)+"</span></div>")
		}else{
			l.GetPosts = append(l.GetPosts, "<div class='post-tagg'><span>#"+item["tag"].(string)+"</span></div>")
		}


		l.GetPosts = append(l.GetPosts, "<div class='post-body'><span>"+item["tips"].(string)+"</span></div>")

		// open post like section
		l.GetPosts = append(l.GetPosts, "<div class='post-like-section'><div class='child-post-like-section'>")

		l.GetPosts = append(l.GetPosts, "<div id='first'>")
		l.GetPosts = append(l.GetPosts, "<span><button id='medal-btn' data-name='"+item["_id"].(primitive.ObjectID).Hex()+"'><div id='medal"+item["_id"].(primitive.ObjectID).Hex()+"' style='pointer-events:none;'>"+CountMedals(item["_id"].(primitive.ObjectID).Hex(), id)+"</div></button></span>")

		if item["thetype"].(string) == "open" {
			l.GetPosts = append(l.GetPosts, "<span><button><i style='pointer-events:none;'  class='fas fa-eye'></i> Public</button></span>")
		}else{
			l.GetPosts = append(l.GetPosts, "<span><button><i style='pointer-events:none;'  class='fas fa-eye'></i> Private</button></span>")
		}

		l.GetPosts = append(l.GetPosts, "<span><button><i style='pointer-events:none;'  class='fas fa-clock'></i> "+item["thetime"].(string)+"</button></span>")

		if item["thetype"].(string) == "open" {
			l.GetPosts = append(l.GetPosts, "<span><button><i style='pointer-events:none;'  class='fa fa-users'></i> "+CountJoin(item["_id"].(primitive.ObjectID).Hex())+"</button></span>")
		}
		l.GetPosts = append(l.GetPosts, "</div>")

		l.GetPosts = append(l.GetPosts, "<div id='second'>")

		if item["thetype"].(string) == "open"{
			if len(FindPublished(item["_id"].(primitive.ObjectID).Hex(), id)) == 0 {
				l.GetPosts = append(l.GetPosts, "<span id='sec-"+item["_id"].(primitive.ObjectID).Hex()+"' "+CheckIfUserJoined(item["_id"].(primitive.ObjectID).Hex(), id)+"><button id='publish-btn' data-name='"+item["_id"].(primitive.ObjectID).Hex()+"'><i style='pointer-events:none;'  class='fas fa-pen'></i> Publish</button></span>")
			}else{
				l.GetPosts = append(l.GetPosts, "<span id='sec-"+item["_id"].(primitive.ObjectID).Hex()+"' "+CheckIfUserJoined(item["_id"].(primitive.ObjectID).Hex(), id)+"><button id='published-btn' data-name='"+item["_id"].(primitive.ObjectID).Hex()+"'><i style='pointer-events:none;'  class='fas fa-pen'></i> Published</button></span>")
			}
		}else{
			l.GetPosts = append(l.GetPosts, "<span><button id='comment-btn' data-name='"+item["_id"].(primitive.ObjectID).Hex()+"'>#Use Tagg</button></span>")
		}
		
		l.GetPosts = append(l.GetPosts, "</div>")

		// close post like section
		l.GetPosts = append(l.GetPosts, "</div></div>")

		//write div
		if item["thetype"].(string) == "open"{
			if len(FindPublished(item["_id"].(primitive.ObjectID).Hex(), id)) == 0 {
				l.GetPosts = append(l.GetPosts, "<div class='post-write-div' id='post-write-div"+item["_id"].(primitive.ObjectID).Hex()+"' "+CheckIfUserJoined(item["_id"].(primitive.ObjectID).Hex(), id)+">")
				l.GetPosts = append(l.GetPosts, "<input type='text' placeholder='Add Challenge' id='writeup"+item["_id"].(primitive.ObjectID).Hex()+"'/>")
				l.GetPosts = append(l.GetPosts, "</div>")
			}
		}

		l.GetPosts = append(l.GetPosts,CheckUserPublished(item["_id"].(primitive.ObjectID).Hex(), id))


		//close post div
		l.GetPosts = append(l.GetPosts, "</div></div>")

	}

}

