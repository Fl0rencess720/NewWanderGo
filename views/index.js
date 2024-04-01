
//注册前先发验证
function registerCaptcha(){
    axios.post("http://localhost:8080/resgister/captcha",{
        user_account : ,
        user_password : ,
    })
    .then(response =>{
        let message = response.data.message;
        console.log(message);
    })
    .catch(error =>{

    })
}
//注册账号
function register(){
    axios.post("http://localhost:8080/resgister",{
        user_account : ,
        user_password : ,
        user_captcha :,
    })
    .then(response =>{
        let message = response.data.message;
        console.log(message);
    })
    .catch(error =>{

    })
}
//登录
function login(){
    axios.post("http://localhost:8080/login",{
        user_account : ,
        user_password : ,
    })
    .then(response =>{
        let message = response.data.message;
        console.log(message);
    })
    .catch(error =>{
    
    })
}
//退出登录
function exit(){
    axios.post("http://localhost:8080/exit")
    .then(response =>{
        let message = response.data.message;
        console.log(message);
    })
    .catch(error =>{

    })
}
//发表评论
function commentsPost() {

    let positionData = {
        x: 115.79941,
        y: 28.656973
    };

    axios.post("http://localhost:8080/comments/post", {
        text: ,
        position: positionData
    })
    .then(response => {
        let placeUID = response.data.place_uid 
        let commentUuid = response.data.comment_uuid;
        let message = response.data.message;
        console.log("Place UID:", placeUID);
        console.log("Comment UUID:", commentUuid);
        console.log("Message:", message);
    })  
    .catch(error => {
        console.error(error);
    });
}
//
function commentsRoam(){
    axios.post("http://localhost:8080/comments/roam", {
        x: ,
        y: ,
    })
    .then(response => {
        let placeUid = response.data.place_uid;
        let commentText = response.data.text;
    
        console.log("Place UID:", placeUid);
        console.log("Comment Text:", commentText);
    })    
    .catch(error => {
        console.error(error);
    });
}
function commentsLike(){
    axios.post("http://localhost:8080/comments/like", {
        comment_uuid = ,
    })
    .then(response => {
        let message = response.data.message;
        console.log(message);
    })    
    .catch(error => {
        console.error(error);
    });
}
function commentsLike(){
    axios.post("http://localhost:8080/places/get", {
        x:,
        y:,
    })
    .then(response => {
        let placeIds = response.data.place_id;
        placeIds.forEach(placeId => {
            console.log("Place ID:", placeId);
        });
    })
    .catch(error => {
        console.error(error);
    });
}
function beginUser(){
    axios.post("http://localhost:8080/begin/user", )
    .then(response => {
        let comments = response.data.comments;
    
        comments.forEach(comment => {
            let userAccount = comment.user_account;
            let date = comment.date;
            let text = comment.text;
            let commentUuid = comment.comment_uuid;
            let placeUid = comment.place_uid;
            let starCount = comment.star_cnt;
    
            console.log("User Account:", userAccount);
            console.log("Date:", date);
            console.log("Text:", text);
            console.log("Comment UUID:", commentUuid);
            console.log("Place UID:", placeUid);
            console.log("Star Count:", starCount);
        });
        let message = response.data.message;
        let userName = response.data.user_name;
    
        console.log("Message:", message);
        console.log("User Name:", userName);
    })
    
    .catch(error => {
        console.error(error);
    });
}

function beginPlaces(){
    axios.post("http://localhost:8080/begin/places", )
    .then(response => {
        let commentsInPlace = response.data.comments_in_place;
        commentsInPlace.forEach(comment => {
            let userAccount = comment.user_account;
            let date = comment.date;
            let text = comment.text;
            let commentUuid = comment.comment_uuid;
            let placeUid = comment.place_uid;
            let starCount = comment.star_cnt;
            console.log("User Account:", userAccount);
            console.log("Date:", date);
            console.log("Text:", text);
            console.log("Comment UUID:", commentUuid);
            console.log("Place UID:", placeUid);
            console.log("Star Count:", starCount);
        });
    })
    
    .catch(error => {
        console.error(error);
    });
}