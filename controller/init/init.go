package Init

import (
	conf "WanderGo/configs"
	user "WanderGo/controller/user"
	mod "WanderGo/models"
	util "WanderGo/utils"
	"log"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

// 载入页面加载个人信息
func LoadPersonalInformation(ctx *gin.Context) {
	//个人评论
	acct := user.SearchAccount(ctx)
	if acct == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "未登录",
		})
		return
	}
	user_ := util.GetUser(acct)
	var u mod.User
	conf.GLOBAL_DB.Preload("Comments").Where("user_account = ?", user_.UserAccount).First(&u)
	//时间排序
	util.NNewComments = u.Comments
	var commentsPayload []user.CommentsPayload
	sort.Sort(util.NNewComments)
	for i := range util.NNewComments {
		commentsPayload = append(commentsPayload, user.CommentsPayload{
			UserAccount: util.NNewComments[i].UserAccount,
			Date:        util.NNewComments[i].Date,
			Text:        util.NNewComments[i].Text,
			PlaceUID:    util.NNewComments[i].PlaceUID,
			CommentUUID: util.NNewComments[i].CommentUUID,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":   "正处于登录状态",
		"comments":  commentsPayload,
		"user_name": user_.UserName,
	})
}
func LoadPlacesInformation(ctx *gin.Context) {
	//地点评论
	var places []mod.Place
	err := conf.GLOBAL_DB.Preload("Comments").Where("is_marked = 1").Find(&places).Error
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "查询地点数据出错",
		})
		return
	}
	for i := range places {
		util.NNewComments = places[i].Comments
		sort.Sort(util.NNewComments)
		places[i].Comments = util.NNewComments
	}
	var comments []mod.Comment
	for i := range places {
		comments = append(comments, places[i].Comments...)
	}
	var commentsPayload []user.CommentsPayload
	for i := range comments {
		commentsPayload = append(commentsPayload, user.CommentsPayload{
			UserAccount: comments[i].UserAccount,
			Date:        comments[i].Date,
			Text:        comments[i].Text,
			PlaceUID:    comments[i].PlaceUID,
			CommentUUID: comments[i].CommentUUID,
		})
	}

	//place_uid是其所在地点的编号
	//comment_uuid是该评论的编号
	ctx.JSON(http.StatusOK, gin.H{
		"comments_in_place": commentsPayload,
	})
}
