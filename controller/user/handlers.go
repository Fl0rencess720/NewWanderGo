package user

import (
	conf "WanderGo/configs"
	mod "WanderGo/models"
	util "WanderGo/utils"
)

func AccountConflictVerification(a string) error { //有错说明不冲突
	var tUser mod.User
	err := conf.GLOBAL_DB.Model(&mod.User{}).Where("user_account = ?", a).First(&tUser).Error
	return err
}
func UserLoginVerification(u mod.User) (int, error) {
	var tUser mod.User
	err := conf.GLOBAL_DB.Model(&mod.User{}).Where(mod.User{UserAccount: u.UserAccount}, "user_account").
		First(&tUser).Error
	if err != nil {
		// 输入账号不存在
		return 1, err
	} else {
		// 若账号存在，检测密码是否正确
		err = conf.GLOBAL_DB.Model(&mod.User{}).Where(mod.User{UserAccount: u.UserAccount, UserPassword: util.EncryptMd5(u.UserPassword)}, "user_account", "user_password").
			First(&tUser).Error
		if err != nil {
			// 密码不正确
			return 2, err
		}
		// 登录成功，没有错误
		return 0, nil
	}
}
