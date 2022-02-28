$(document).ready(function(){
//     $("#join-btn").click(function(){
//       $.post("/edit/join/",
//         {
//           postid: $("#join-btn").attr("data-name")
//         },
//         function(data,status){
//             if(status === "success"){
//                 if(data === "inserted"){
//                     $("#join-btn"+$("#join-btn").attr("data-name")).html("Joined")
//                     $("#post-write-div"+$("#join-btn").attr("data-name")).show()
//                     $("#sec-"+$("#join-btn").attr("data-name")).show()
//                 }else{
//                     $("#join-btn"+$("#join-btn").attr("data-name")).html("Join")
//                     $("#post-write-div"+$("#join-btn").attr("data-name")).hide()
//                     $("#sec-"+$("#join-btn").attr("data-name")).hide()
//                 }
//             }
//         });
//     });

   

    $(document).click(function(event){

        if(event.target.id == "medal-btn"){
            $.post("/edit/medal/",
            {
                postid: event.target.getAttribute('data-name')
            },
            
            function(data,status){
                if(status === "success"){
                    $("#medal"+event.target.getAttribute('data-name')).html(data)
                }
            });  
        }
    })

    $(document).click(function(event){

        if(event.target.id == "publish-btn"){

            if ($("#writeup"+event.target.getAttribute('data-name')).val().length < 50){
                alert("Write up not up too 50 characters")
            }else{
                $.post("/edit/publish/",
                {
                  postid: $("#publish-btn").attr("data-name"),
                  data:$("#writeup"+event.target.getAttribute('data-name')).val()
                },
                function(data,status){
                    if(status === "success"){
                        $("#writeup"+event.target.getAttribute('data-name')).hide()
                        alert("done")
                    }
                });
            }
 
        }else if (event.target.id == "btn-challengers"){
            alert("webbee")
        }
    })


    // $("#publish-btn").click(function(){
    //     if($("#writeup"+$("#publish-btn").attr("data-name")).val().length > 0 ){
    //         $.post("/edit/publish/",
    //         {
    //           postid: $("#publish-btn").attr("data-name"),
    //           data:$("#writeup"+$("#publish-btn").attr("data-name")).val()
    //         },
    //         function(data,status){
    //             if(status === "success"){
    //                 // $("#medal"+$("#medal-btn").attr("data-name")).html(data)
    //                 alert("done")
    //             }
    //         });
    //     }else{
    //         alert("Cant publish empty write up")
    //     }
    // });
})
