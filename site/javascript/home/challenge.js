$(document).ready(function(){
    $(document).click(function(event){
        if(event.target.id == "a"){
            $("#the-option").html("Option <b>A</b>")
            $("#option-collect").val("a")
            $("#support-box").show()
        }else if(event.target.id == "b"){
            $("#the-option").html("Option <b>B</b>")
            $("#option-collect").val("b")
            $("#support-box").show()
        }else if(event.target.id == "c"){
            $("#the-option").html("Option <b>C</b>")
            $("#option-collect").val("c")
            $("#support-box").show()
        }
    })

    $(document).click(function(event){
        if(event.target.id == "aa"){
            $.post("/edit/choice/",
            {
                postid: event.target.getAttribute('data-name'),
                choice:"a"
            },
            
            function(data,status){
                if(status === "success"){
                    $("#the-challenge-q").html(data)
                }
            });

        }else  if(event.target.id == "bb"){
            $.post("/edit/choice/",
            {
                postid: event.target.getAttribute('data-name'),
                choice:"b"
            },
            
            function(data,status){
                if(status === "success"){
                    $("#the-challenge-q").html(data)
                }
            });
        }else  if(event.target.id == "cc"){
            $.post("/edit/choice/",
            {
                postid: event.target.getAttribute('data-name'),
                choice:"c"
            },
            
            function(data,status){
                if(status === "success"){
                    $("#the-challenge-q").html(data)
                }
            });
        }
    })

})