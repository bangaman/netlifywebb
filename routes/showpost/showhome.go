package showpost

import (
	"mungo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "fmt"
	// "strings"
)


func (l *ShowPostStruct) HomeQuestion(id string){

	getc := mungo.Connect("app").Collection("questions")
	stuff := mungo.QueryJoin(getc)
	// fmt.Println(stuff)

	for _, item := range stuff {

		// open challenge link
		l.HomePosts = append(l.HomePosts, "<a href='/challenge/"+item["_id"].(primitive.ObjectID).Hex()+"' style='text-decoration:none;color:black;'>")

		l.HomePosts = append(l.HomePosts, "<div class='post-section' style='margin-top:15px;'><div class='child-post-section' style='border:1px solid #ccc;'>")

		//open top details div
		l.HomePosts = append(l.HomePosts, "<div class='post-top-details'>")

		l.HomePosts = append(l.HomePosts,  "<div class='post-section-image'>")
		l.HomePosts = append(l.HomePosts, "<div id='post-pic'><button>"+item["theid"].(primitive.A)[0].(primitive.M)["username"].(string)[0:1]+"</button></div>")
		l.HomePosts = append(l.HomePosts, "<div id='post-username'><span>"+item["theid"].(primitive.A)[0].(primitive.M)["username"].(string)+"</span><div id='category-select'>"+item["subject"].(string)+"</div></div>")
		l.HomePosts = append(l.HomePosts, "<div id='post-date'><span>"+item["time"].(string)+"</span></div>")
		l.HomePosts = append(l.HomePosts, "</div>")
		
		//close top details div
		l.HomePosts = append(l.HomePosts, "</div>")

		l.HomePosts = append(l.HomePosts,  "<div class='post-body'><span>"+item["q"].(string)+"</span></div>")

		l.HomePosts = append(l.HomePosts,  "<div class='available' style='margin-top:5px;'><span>available options <b>A</b>, <b>B</b>, <b>C</b> and <b>D</b></span></div>")

		// open post like section
		l.HomePosts = append(l.HomePosts, "<div class='post-like-section'><div class='child-post-like-section'>")

		l.HomePosts = append(l.HomePosts, "<div id='first'>")
		l.HomePosts = append(l.HomePosts, "<span><button id='medal-btn'><div style='pointer-events:none;'>"+CountMedals(item["_id"].(primitive.ObjectID).Hex(), id)+"</div></button></span>")

		l.HomePosts = append(l.HomePosts, "<span><button><i style='pointer-events:none;'  class='fas fa-eye'></i> Public</button></span>")
		
		// l.HomePosts = append(l.HomePosts, "<span><button><i style='pointer-events:none;'  class='fas fa-clock'></i> "+item["time"].(string)+"</button></span>")
		l.HomePosts = append(l.HomePosts, "</div>")

	
		l.HomePosts = append(l.HomePosts, "<div id='second'>")
		l.HomePosts = append(l.HomePosts, "<span><button><i class='fas fa-pen'></i> Answer Question</button></span>")
		l.HomePosts = append(l.HomePosts, "</div>")
		

		// close post like section
		l.HomePosts = append(l.HomePosts, "</div></div>")


		//close post div
		l.HomePosts = append(l.HomePosts, "</div></div>")

		// close challenge link
		l.HomePosts = append(l.HomePosts, "</a>")
	}
}




func (l *ShowPostStruct) HomeChallenge(id string){

	getc := mungo.Connect("app").Collection("post")
	stuff := mungo.QueryJoin(getc)

	for _, item := range stuff {

		// open challenge link
		l.HomePosts = append(l.HomePosts, "<a href='/challenge/"+item["_id"].(primitive.ObjectID).Hex()+"' style='text-decoration:none;color:black;'>")

		l.HomePosts = append(l.HomePosts, "<div class='post-section' style='margin-top:15px;'><div class='child-post-section' style='border:1px solid #ccc;'>")

		//open top details div
		l.HomePosts = append(l.HomePosts, "<div class='post-top-details'>")

		l.HomePosts = append(l.HomePosts,  "<div class='post-section-image'>")
		l.HomePosts = append(l.HomePosts, "<div id='post-pic'><button>"+item["theid"].(primitive.A)[0].(primitive.M)["username"].(string)[0:1]+"</button></div>")
		l.HomePosts = append(l.HomePosts, "<div id='post-username'><span>"+item["theid"].(primitive.A)[0].(primitive.M)["username"].(string)+"</span><div id='category-select'>"+item["category"].(string)+"</div></div>")
		l.HomePosts = append(l.HomePosts, "<div id='post-date'><span>"+item["time"].(string)+"</span></div>")
		l.HomePosts = append(l.HomePosts, "</div>")
		
		//close top details div
		l.HomePosts = append(l.HomePosts, "</div>")

		l.HomePosts = append(l.HomePosts,  "<div class='post-body'><span>"+item["tips"].(string)+"</span></div>")

		// open post like section
		l.HomePosts = append(l.HomePosts, "<div class='post-like-section'><div class='child-post-like-section'>")

		l.HomePosts = append(l.HomePosts, "<div id='first'>")
		l.HomePosts = append(l.HomePosts, "<span><button id='medal-btn'><div style='pointer-events:none;'>"+CountMedals(item["_id"].(primitive.ObjectID).Hex(), id)+"</div></button></span>")

		if item["thetype"].(string) == "open" {
			l.HomePosts = append(l.HomePosts, "<span><button><i style='pointer-events:none;'  class='fas fa-eye'></i> Public</button></span>")
		}else{
			l.HomePosts = append(l.HomePosts, "<span><button><i style='pointer-events:none;'  class='fas fa-eye'></i> Private</button></span>")
		}

		l.HomePosts = append(l.HomePosts, "<span><button><i style='pointer-events:none;'  class='fas fa-clock'></i> "+item["thetime"].(string)+"</button></span>")

		if item["thetype"].(string) == "open" {
			l.HomePosts = append(l.HomePosts, "<span><button><i style='pointer-events:none;'  class='fa fa-users'></i> "+CountJoin(item["_id"].(primitive.ObjectID).Hex())+"</button></span>")
		}
		l.HomePosts = append(l.HomePosts, "</div>")

		if item["thetype"].(string) == "open" {
			l.HomePosts = append(l.HomePosts, "<div id='second'>")
			l.HomePosts = append(l.HomePosts, "<span><button><i class='fas fa-pen'></i> Start Challenge</button></span>")
			l.HomePosts = append(l.HomePosts, "</div>")
		}

		// close post like section
		l.HomePosts = append(l.HomePosts, "</div></div>")

		//close post div
		l.HomePosts = append(l.HomePosts, "</div></div>")

		// close challenge link
		l.HomePosts = append(l.HomePosts, "</a>")
	}
}