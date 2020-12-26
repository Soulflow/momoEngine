package core

import (
	"github.com/dgrijalva/jwt-go"
	"momoEngine/utils"
	"time"
)

/**
 * 检查token
 * @param token
 */
func CheckToken(token string) bool {
	result, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return []byte(utils.GetConfig().JwtSecret), nil
	})
	if err != nil {
		if err.Error() == "Token is expired" && result != nil {
			utils.Info("[JWT] token is expired")
		} else {
			utils.Warn("[JWT] parse with claims failed")
		}
		return false
	}
	return true
}

/**
 * 创建token
 * @param uid 用户id
 */
func CreateToken(uid string) string {
	claims := &jwt.StandardClaims{
		NotBefore: time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Issuer:    "r3inb",
		Id:        uid,
	}
	utils.Info("[JWT] create token -> " + uid)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(utils.GetConfig().JwtSecret))
	if err != nil {
		println(err.Error())
		return ""
	}
	return ss
}

/**
 * 获取用户id
 */
func getIdByToken(token string) string {
	result, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return []byte(utils.GetConfig().JwtSecret), nil
	})
	if err != nil {
		if err.Error() == "Token is expired" && result != nil {
			utils.Info("token is expired " + result.Claims.(jwt.MapClaims)["jti"].(string))
		} else {
			utils.Warn("parse with claims failed")
		}
		return ""
	}
	return result.Claims.(jwt.MapClaims)["jti"].(string)
}
