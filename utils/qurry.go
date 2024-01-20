package Util

import (
	conf "WanderGo/configs"
	mod "WanderGo/models"
	"log"
)

// 用user_account从数据库中找到user
func GetUser(a string) mod.User {
	var u mod.User
	err := conf.GLOBAL_DB.Model(&mod.User{}).Where("user_account = ?", a).First(&u).Error
	if err != nil {
		log.Panicln(err)
		return u
	}
	return u
}
