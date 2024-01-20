package user

import (
	conf "WanderGo/configs"
	pos "WanderGo/controller/position"
	mod "WanderGo/models"
	util "WanderGo/utils"
	"fmt"

	"log"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func PostComment(ctx *gin.Context) {
	var com mod.Comment
	err := ctx.ShouldBind(&com)
	if err != nil {
		log.Println(err)
		return
	}
	centerPoint := pos.PositionHandlerComment(com.Position)
	fmt.Println(centerPoint)
	place := pos.GetPos(centerPoint)
	fmt.Printf("place: %v\n", place)
	com.UserAccount = SearchAccount(ctx)
	user := util.GetUser(com.UserAccount)
	com.User = user
	com.Place = place
	date := time.Now().Format("2006-01-02 15:04:05")
	com.Date = date
	currentTime := strconv.FormatInt(time.Now().Unix(), 10)
	com.CommentUUID = util.EncryptMd5(user.UserAccount + currentTime)
	if com.Text != "" || com.PhotoData != nil {
		err := conf.GLOBAL_DB.Model(&mod.Comment{}).Create(&com).Error
		if err != nil {
			log.Println(err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message":      "成功发布评论",
			"comment_uid":  place.ID,
			"comment_uuid": com.CommentUUID,
		})
	}
}
func GetAccountWithComments(accountID uint) (mod.User, error) {
	var u mod.User
	err := conf.GLOBAL_DB.Preload("Comment").First(&mod.User{}, accountID).Error
	if err != nil {
		return mod.User{}, err
	}
	return u, nil
}
func GetPlaceWithComments(placeID uint) (mod.Place, error) {
	var place mod.Place
	err := conf.GLOBAL_DB.Preload("Comment").First(&place, placeID).Error
	if err != nil {
		return mod.Place{}, err
	}
	return place, nil
}
func TestComments(ctx *gin.Context) {
	// var com dbf.User
	// dbf.GLOBAL_DB.Preload("Comments").Take(&com)
	// fmt.Println(com)
	var p mod.User
	conf.GLOBAL_DB.Preload("Comments").Where("user_account = ?", "panyaan@ncuhome.club").First(&p)
	// con.GLOBAL_DB.Preload("Comments").Take(&p)
	// var pp con.Place
	// con.GLOBAL_DB.Preload("Comments").Find(&pp)
	fmt.Println(p)
}
func GetComment(c string) mod.Comment {
	var com mod.Comment
	err := conf.GLOBAL_DB.Model(&mod.Comment{}).Where("comment_uuid = ?", c).First(&com).Error
	if err != nil {
		log.Println(err)
		return mod.Comment{}
	}
	return com
}

// 时间排序
func HandleNewComments(ctx *gin.Context) {
	err := conf.GLOBAL_DB.Model(&mod.Comment{}).Find(&util.NNewComments).Error
	if err != nil {
		log.Println(err)
		return
	}
	sort.Sort(util.NNewComments)
	ctx.JSON(http.StatusOK, util.NNewComments)
}

// 点赞数排序
func HandleHotComments(ctx *gin.Context) {
	err := conf.GLOBAL_DB.Model(&mod.Comment{}).Find(&util.HHotComments).Error
	if err != nil {
		log.Println(err)
		return
	}
	sort.Sort(util.HHotComments)
	ctx.JSON(http.StatusOK, util.HHotComments)
}
