package Position

import (
	conf "WanderGo/configs"
	mod "WanderGo/models"
	util "WanderGo/utils"
	"log"
	"math/rand"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 前端传入坐标,返回地区中心点(已弃用)
// func PositionHandler(ctx *gin.Context) mod.Address {
// 	var pos mod.Address
// 	err := ctx.ShouldBind(&pos)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	var place mod.Place
// 	err = conf.GLOBAL_DB.Model(&mod.Place{}).Where("JSON_EXTRACT(top_left_point, '$.x') <= ? AND JSON_EXTRACT(top_left_point, '$.y') >= ? AND JSON_EXTRACT(bottom_right_point, '$.x') >= ? AND JSON_EXTRACT(bottom_right_point, '$.y') <= ?", pos.X, pos.Y, pos.X, pos.Y).First(&place).Error
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	return place.CenterPoint
// }

// 前端传入坐标，返回中心点在250m内的地点名(已弃用)
// func PositionsHandler(ctx *gin.Context) {
// 	var pos mod.Address
// 	err := ctx.ShouldBind(&pos)
// 	if err != nil {
// 		log.Println(err)

// 	}
// 	// var places []conf.Place
// 	// conf.GLOBAL_DB.Model(&conf.Place{}).Where("(POWER(? - JSON_EXTRACT(center_point, '$.x'), 2) / POWER(0.0028525, 2)) + (POWER(? - JSON_EXTRACT(center_point, '$.y'), 2) / POWER(0.0022475, 2)) <= 1", pos.X, pos.Y).Find(&places)
// 	// var plIDs []int
// 	// for i := range places {
// 	// 	plIDs = append(plIDs, int(places[i].ID))
// 	// }
// 	res, err := conf.GLOBAL_RDB.GeoRadius("NCU:Buildings", pos.X, pos.Y, &redis.GeoRadiusQuery{Radius: 250, Unit: "m"}).Result()
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"places": res,
// 	})

// }
func FindPlaceByPoint(db *gorm.DB, pointX, pointY float64) (mod.Place, error) {
	var place mod.Place
	if err := db.Raw(`
        SELECT place_name, ST_Astext(geo_info) AS geo_info
        FROM place_test
        WHERE ST_CONTAINS(geo_info, POINT(?, ?))
    `, pointX, pointY).Scan(&place).Error; err != nil {
		return mod.Place{}, err
	}
	return place, nil
}
func PositionsHandler(ctx *gin.Context) {
	var pos mod.Address
	err := ctx.ShouldBindJSON(&pos)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	placeTest, err := FindPlaceByPoint(conf.GLOBAL_DB, pos.X, pos.Y)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"place_id":   placeTest.ID,
		"place_name": placeTest.PlaceName})
}

// 以上两种方法待测试完精确度后再做选择
func GetPos(c mod.Address) mod.Place {
	var p mod.Place
	err := conf.GLOBAL_DB.Model(&mod.Place{}).Where("JSON_EXTRACT(center_point, '$.x') = ? AND JSON_EXTRACT(center_point, '$.y') = ?", c.X, c.Y).First(&p).Error
	if err != nil {
		log.Println(err)
	}
	return p
}

// func PositionHandlerComment(a mod.Address) mod.Address {
// 	var place mod.Place
// 	err := conf.GLOBAL_DB.Model(&mod.Place{}).Where("JSON_EXTRACT(top_left_point, '$.x') <= ? AND JSON_EXTRACT(top_left_point, '$.y') >= ? AND JSON_EXTRACT(bottom_right_point, '$.x') >= ? AND JSON_EXTRACT(bottom_right_point, '$.y') <= ?", a.X, a.Y, a.X, a.Y).First(&place).Error
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	return place.CenterPoint
// }

// 随机漫游?400m
func Roaming(ctx *gin.Context) {
	var pos mod.Address
	if err := ctx.ShouldBind(&pos); err != nil {
		log.Println(err)
		return
	}
	var places []mod.Place
	conf.GLOBAL_DB.Model(&mod.Place{}).
		Where(`ST_Distance(center_point, ST_GeomFromText('POINT(? ?)')) <= 400`, pos.X, pos.Y).
		Find(&places)
	randomIndex := rand.Intn(len(places))
	selectedPlace := places[randomIndex]
	var pl mod.Place
	conf.GLOBAL_DB.Preload("Comments").Where("id = ?", selectedPlace.ID).First(&pl)
	if len(pl.Comments) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message":    "此处没有漫游点",
			"place_name": selectedPlace.PlaceName,
		})
		return
	}
	randomCommentIndex := rand.Intn(4)
	util.HHotComments = pl.Comments
	sort.Sort(util.HHotComments)
	selectedComment := util.HHotComments[randomCommentIndex]
	ctx.JSON(http.StatusOK, gin.H{
		"place_uid":    selectedComment.PlaceUID,
		"comment_uuid": selectedComment.CommentUUID,
		"text":         selectedComment.Text,
	})
	ctx.Data(http.StatusOK, "image/jpeg", selectedComment.PhotoData)
}

// 点亮地区,前端满足条件该place:{IsMarked:false}且定位到该place时发送请求，
// func MarkPlace(ctx *gin.Context) {
// 	var place mod.Place
// 	err := ctx.ShouldBind(&place)
// 	if err != nil {
// 		ctx.JSON(http.StatusOK, gin.H{"message": "参数错误,绑定失败"})
// 		return
// 	}
// 	if err := conf.GLOBAL_DB.Where("id = ?", place.ID).First(&place).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			ctx.JSON(http.StatusNotFound, gin.H{"message": "未找到对应的地区"})
// 			return
// 		}
// 		log.Println("Error querying database:", err)
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "查询数据库失败"})
// 		return
// 	}
// 	if place.IsMarked {
// 		ctx.JSON(http.StatusOK, gin.H{"message": "该地区已被点亮过"})
// 		return
// 	}
// 	// 设置 IsMarked 为 true，并保存到数据库
// 	place.IsMarked = true
// 	if err := conf.GLOBAL_DB.Save(&place).Error; err != nil {
// 		log.Println("Error updating place:", err)
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "更新至数据库失败"})
// 		return
// 	}
// 	// 返回标记成功信息
// 	ctx.JSON(http.StatusOK, gin.H{
// 		place.PlaceName: "已被点亮",
// 	})
// }
