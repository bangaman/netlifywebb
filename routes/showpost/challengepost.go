package showpost

import (
	"mungo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "fmt"
	// "strings"
)



func GetAllPosts(postid, id string) []string {

	getc := mungo.Connect("app").Collection("post")
	stuff := mungo.QueryJoin(getc)

	getposts := []string{}

	for _, item := range stuff {

		if item["_id"].(primitive.ObjectID).Hex() == postid {

			getposts = append(getposts, "<div class='post-section' style='margin-top:15px;'><div class='child-post-section' style='border:1px solid #ccc;'>")
			
			//open top details div
			getposts = append(getposts, "<div class='post-top-details'>")

			getposts = append(getposts,  "<div class='post-section-image'>")
			getposts = append(getposts, "<div id='post-pic'><button>"+item["theid"].(primitive.A)[0].(primitive.M)["username"].(string)[0:1]+"</button></div>")
			getposts = append(getposts, "<div id='post-username'><span>"+item["theid"].(primitive.A)[0].(primitive.M)["username"].(string)+"</span><div id='category-select'>"+item["category"].(string)+"</div></div>")
		    getposts = append(getposts, "<div id='post-date'><span>"+item["time"].(string)+"</span></div>")
			getposts = append(getposts, "</div>")
			
			// if item["_id"].(primitive.ObjectID).Hex() 
			if item["theid"].(primitive.A)[0].(primitive.M)["_id"].(primitive.ObjectID).Hex() != id {
				getposts = append(getposts,  "<div class='post-section-join'>")
				getposts = append(getposts,"<form method='post' action='/edit/join/'><button id='join-btn' data-name='"+item["_id"].(primitive.ObjectID).Hex()+"'><span id='join-btn"+item["_id"].(primitive.ObjectID).Hex()+"' style='pointer-events:none;'>"+FindJoin(item["_id"].(primitive.ObjectID).Hex(), id)+"</span></button>")
				getposts = append(getposts, "<input style='display:none;' type='text' name='join' value='"+item["_id"].(primitive.ObjectID).Hex()+"'/></form>")
				getposts = append(getposts, "</div>")
			}
			
			//close top-details div
			getposts = append(getposts, "</div>")

			getposts = append(getposts, "<div class='post-body-title'><h3>Writeup Challenge</h3></div>")
			getposts = append(getposts,  "<div class='post-body'><span>"+item["tips"].(string)+"</span></div>")

			// open post like section
			getposts = append(getposts, "<div class='post-like-section'><div class='child-post-like-section'>")

			getposts = append(getposts, "<div id='first'>")
		
			getposts = append(getposts, "<span><button id='medal-btn' data-name='"+item["_id"].(primitive.ObjectID).Hex()+"'><div id='medal"+item["_id"].(primitive.ObjectID).Hex()+"' style='pointer-events:none;'>"+CountMedals(item["_id"].(primitive.ObjectID).Hex(), id)+"</div></button></span>")

			if item["thetype"].(string) == "open" {
				getposts = append(getposts, "<span><button><i style='pointer-events:none;'  class='fas fa-eye'></i> Public</button></span>")
			}else{
				getposts = append(getposts, "<span><button><i style='pointer-events:none;'  class='fas fa-eye'></i> Private</button></span>")
			}
			getposts = append(getposts, "<span><button><i style='pointer-events:none;'  class='fas fa-clock'></i> "+item["thetime"].(string)+"</button></span>")
		
			if item["thetype"].(string) == "open" {
				getposts = append(getposts, "<span><button><i style='pointer-events:none;'  class='fa fa-users'></i> "+CountJoin(item["_id"].(primitive.ObjectID).Hex())+"</button></span>")
			}
			getposts = append(getposts, "</div>")
		

			if item["thetype"].(string) == "open"{
				getposts = append(getposts, "<div id='second'>")

				if len(FindPublished(item["_id"].(primitive.ObjectID).Hex(), id)) == 0 {
					getposts = append(getposts, "<span id='sec-"+item["_id"].(primitive.ObjectID).Hex()+"' "+CheckIfUserJoined(item["_id"].(primitive.ObjectID).Hex(), id)+"><button id='publish-btn' data-name='"+item["_id"].(primitive.ObjectID).Hex()+"'> Publish</button></span>")
				}else{
					getposts = append(getposts,"<span id='sec-"+item["_id"].(primitive.ObjectID).Hex()+"' "+CheckIfUserJoined(item["_id"].(primitive.ObjectID).Hex(), id)+"><button id='published-btn' data-name='"+item["_id"].(primitive.ObjectID).Hex()+"'><i style='pointer-events:none;'  class='fas fa-pen'></i> Published</button></span>")			
				}

				getposts = append(getposts, "</div>")

			}
		
			// close post like section
			getposts = append(getposts, "</div></div>")

			if FindJoin(item["_id"].(primitive.ObjectID).Hex(), id) == "Joined" {
				if len(FindPublished(item["_id"].(primitive.ObjectID).Hex(), id)) == 0 {
					getposts = append(getposts, "<div class='write-up'>")
					
					// tools
					getposts = append(getposts, "<textarea id='writeup"+item["_id"].(primitive.ObjectID).Hex()+"'></textarea>")
					//tools
					
					getposts = append(getposts, "</div>")
				}
			}

			getposts = append(getposts, CheckUserPublished(item["_id"].(primitive.ObjectID).Hex(), id))

			getposts = append(getposts, "</div></div>")
		}
	}
	
	return getposts

}



func GetAllQuestions(postid, id string) []string {

	getc := mungo.Connect("app").Collection("questions")
	stuff := mungo.QueryJoin(getc)

	getposts := []string{}

	for _, item := range stuff {

		if item["_id"].(primitive.ObjectID).Hex() == postid {

			getposts = append(getposts, "<div class='post-section' style='margin-top:15px;'><div class='child-post-section' style='border:1px solid #ccc;'>")
			
			//open top details div
			getposts = append(getposts, "<div class='post-top-details'>")

			getposts = append(getposts,  "<div class='post-section-image'>")
			getposts = append(getposts, "<div id='post-pic'><button>"+item["theid"].(primitive.A)[0].(primitive.M)["username"].(string)[0:1]+"</button></div>")
			getposts = append(getposts, "<div id='post-username'><span>"+item["theid"].(primitive.A)[0].(primitive.M)["username"].(string)+"</span><div id='category-select'>"+item["subject"].(string)+"</div></div>")
		    getposts = append(getposts, "<div id='post-date'><span>"+item["time"].(string)+"</span></div>")
			getposts = append(getposts, "</div>")
			
			//close top-details div
			getposts = append(getposts, "</div>")

			getposts = append(getposts, "<div class='post-body-title'><h3>Question Challenge</h3></div>")
			getposts = append(getposts,  "<div class='post-body'><span>"+item["q"].(string)+"</span></div>")

			if item["check"].(string) == "on"{

				if CheckUserAnsweredQuestionOrNot(item["_id"].(primitive.ObjectID).Hex(), id) == "true"{
					//options
					getposts = append(getposts, "<div class='options'>")
					getposts = append(getposts, "<span><button><b>A</b> "+item["a"].(string)+"</button></span>")
					getposts = append(getposts, "<span><button><b>B</b> "+item["b"].(string)+"</button></span>")
					getposts = append(getposts, "<span><button ><b>C</b> "+item["c"].(string)+"</button></span>")
					getposts = append(getposts, "<span><button><b>D</b> "+item["d"].(string)+"</button></span>")
					getposts = append(getposts,  "</div>")
					//options
				}else{
					//options
					getposts = append(getposts, "<div class='options'>")
					getposts = append(getposts, "<span><button id='a'><b>A</b> "+item["a"].(string)+"</button></span>")
					getposts = append(getposts, "<span><button id='b'><b>B</b> "+item["b"].(string)+"</button></span>")
					getposts = append(getposts, "<span><button id='c'><b>C</b> "+item["c"].(string)+"</button></span>")
					getposts = append(getposts, "<span><button id='d'><b>D</b> "+item["d"].(string)+"</button></span>")
					getposts = append(getposts,  "</div>")
					//options
				}
				
			}else{
				//options
				getposts = append(getposts, "<div class='options'>")

				if FindChoice(item["_id"].(primitive.ObjectID).Hex(), id)[1:] == "T" {

					if item["default"] != nil {
						if FindChoice(item["_id"].(primitive.ObjectID).Hex(), id)[0:1] == item["default"].(string) {
							getposts = append(getposts, "<div class='default'><span><b>Correct</b> You are genius</span></div>")
						} else{
							getposts = append(getposts, "<div class='default'><span><b>Wrong</b> the answer is <b style='text-transform:uppercase;'>"+item["default"].(string)+"</b></span></div>")
						}
					}
					if FindChoice(item["_id"].(primitive.ObjectID).Hex(), id)[0:1] == "a"{
						getposts = append(getposts, "<span><button style='background:cadetblue;color:white;font-weight:900;' id='aee' data-name='"+item["_id"].(primitive.ObjectID).Hex()+"'><b>A</b> "+item["a"].(string)+"</button></span>")
					}else{
						getposts = append(getposts, "<span><button id='aee' data-name='"+item["_id"].(primitive.ObjectID).Hex()+"'><b>A</b> "+item["a"].(string)+"</button></span>")
					}

					if FindChoice(item["_id"].(primitive.ObjectID).Hex(), id)[0:1] == "b"{
						getposts = append(getposts, "<span><button style='background:cadetblue;color:white;font-weight:900;' id='bee' data-name='"+item["_id"].(primitive.ObjectID).Hex()+"'><b>B</b> "+item["b"].(string)+"</button></span>")
					}else{
						getposts = append(getposts, "<span><button id='bee' data-name='"+item["_id"].(primitive.ObjectID).Hex()+"'><b>B</b> "+item["b"].(string)+"</button></span>")
					}

					if FindChoice(item["_id"].(primitive.ObjectID).Hex(), id)[0:1] == "c"{
						getposts = append(getposts, "<span><button style='background:cadetblue;color:white;font-weight:900;' id='cee' data-name='"+item["_id"].(primitive.ObjectID).Hex()+"'><b>C</b> "+item["c"].(string)+"</button></span>")
					}else{
						getposts = append(getposts, "<span><button id='cee' data-name='"+item["_id"].(primitive.ObjectID).Hex()+"'><b>C</b> "+item["c"].(string)+"</button></span>")
					}

					if FindChoice(item["_id"].(primitive.ObjectID).Hex(), id)[0:1] == "d"{
						getposts = append(getposts, "<span><button style='background:cadetblue;color:white;font-weight:900;' id='dee' data-name='"+item["_id"].(primitive.ObjectID).Hex()+"'><b>D</b> "+item["d"].(string)+"</button></span>")
					}else{
						getposts = append(getposts, "<span><button id='dee' data-name='"+item["_id"].(primitive.ObjectID).Hex()+"'><b>D</b> "+item["d"].(string)+"</button></span>")
					}
				}else{
					getposts = append(getposts, "<span><button id='aa' data-name='"+item["_id"].(primitive.ObjectID).Hex()+"'><b>A</b> "+item["a"].(string)+"</button></span>")
					getposts = append(getposts, "<span><button id='bb' data-name='"+item["_id"].(primitive.ObjectID).Hex()+"'><b>B</b> "+item["b"].(string)+"</button></span>")
					getposts = append(getposts, "<span><button id='cc' data-name='"+item["_id"].(primitive.ObjectID).Hex()+"'><b>C</b> "+item["c"].(string)+"</button></span>")
					getposts = append(getposts, "<span><button id='dd' data-name='"+item["_id"].(primitive.ObjectID).Hex()+"'><b>D</b> "+item["d"].(string)+"</button></span>")
				}
				
				getposts = append(getposts,  "</div>")
				//options
			}

			// open post like section
			getposts = append(getposts, "<div class='post-like-section'><div class='child-post-like-section'>")

			getposts = append(getposts, "<div id='first'>")
		
			getposts = append(getposts, "<span><button id='medal-btn' data-name='"+item["_id"].(primitive.ObjectID).Hex()+"'><div id='medal"+item["_id"].(primitive.ObjectID).Hex()+"' style='pointer-events:none;'>"+CountMedals(item["_id"].(primitive.ObjectID).Hex(), id)+"</div></button></span>")

			getposts = append(getposts, "<span><button><i style='pointer-events:none;'  class='fas fa-eye'></i> Public</button></span>")
			
		
			// close post like section
			getposts = append(getposts, "</div></div>")

			if item["check"].(string) == "on"{
				//suport
				getposts = append(getposts, "<div class='support' id='support-box' style='display:none;'>")
				getposts = append(getposts, "<form method='post' action='/edit/q/'>")
				getposts = append(getposts, "<div id='the-option'></div>")
				getposts = append(getposts, "<textarea name='support' placeholder='Support your answer' required></textarea>")
				getposts = append(getposts, "<input type='text' style='display:none;' id='option-collect' name='option'>")
				getposts = append(getposts, "<input type='text' style='display:none;' name='postid' value='"+item["_id"].(primitive.ObjectID).Hex()+"'>")
				getposts = append(getposts, "<span><button>Answer</button></span>")
				getposts = append(getposts, "</form>")
				getposts = append(getposts, "</div>")
				// support

			} 
			if item["default"] != nil {
				getposts = append(getposts, CheckUserAnsweredQuestion(item["_id"].(primitive.ObjectID).Hex(),item["default"].(string), id))
			}else{
				getposts = append(getposts, CheckUserAnsweredQuestion(item["_id"].(primitive.ObjectID).Hex(),"", id))
			}

			getposts = append(getposts, "</div></div>")

		}
	}
	
	return getposts

}