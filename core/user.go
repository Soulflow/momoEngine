package core

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"momoEngine/utils"
)

type User struct {
	Uid      string `json:"uid" gorm:"primary_key" valid:"alphanum,length(5|12)" `
	Username string `json:"username" valid:"required,length(3|16)"`
	Password string `json:"password,omitempty" valid:"ascii,length(3|16)"`
	Sex      int    `json:"sex"`
	Role     string `json:"role"`
	Avatar   string `json:"avatar"`
	Status   int    `json:"status"`
}

type UserInfo struct {
	Uid      string `json:"uid" valid:"alphanum,length(5|12)"`
	Password string `json:"password" valid:"ascii,length(3|16)"`
}

/**
 * 登录用户
 * @param user UserInfo
 * @return string token
 * @param error login info
 */
func LoginUser(user UserInfo) (string, error) {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return "", err
	}

	if userDb, exist := FoundUser(user.Uid); exist {
		return "", errors.New("non existing user")
	} else {
		// 状态判断
		if userDb.Status == utils.StatusLock {
			return "", errors.New("your account has been locked by system")
		} else if userDb.Status == utils.StatusAudit {
			return "", errors.New("your account is waiting for audit")
		}
		if userDb.Password == utils.CreateMD5(user.Password) {
			return CreateToken(userDb.Uid), nil
		}
		return "", errors.New("user password error")
	}
}

/**
 * 用户是否存在
 * @param uid 用户id
 */
func FoundUser(uid string) (User, bool) {
	var user User
	if GetDB().Where("uid = ?", uid).First(&user).RecordNotFound() {
		return user, true
	}
	return user, false
}
