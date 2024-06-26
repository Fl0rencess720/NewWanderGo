package user

import (
	conf "WanderGo/configs"
	mod "WanderGo/models"
	util "WanderGo/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LikeHandler(ctx *gin.Context) { //"comment_uuid"
	var star mod.Star
	err := ctx.ShouldBind(&star)
	if err != nil {
		log.Println(err)
		return
	}
	star.UserAccount = SearchAccount(ctx)
	com := GetComment(star.CommentUUID)
	fmt.Println(com)
	user := util.GetUser(star.UserAccount)
	star.User = user
	star.Comment = com
	err = conf.GLOBAL_DB.Model(&mod.Star{}).Create(&star).Error
	if err != nil {
		log.Println(err)
		return
	}
	com.StarCnt++
	fmt.Println(com.StarCnt)
	err = conf.GLOBAL_DB.Model(&mod.Comment{}).Where("id = ?", com.ID).Select("star_cnt").Updates(mod.Comment{StarCnt: com.StarCnt}).Error
	if err != nil {
		log.Panicln(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "点赞成功",
	})
}
