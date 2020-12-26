package core

import (
	"momoEngine/utils"
	"testing"
)

func TestCheckToken(t *testing.T) {
	println(CheckToken("1213123"))
}

func TestCreateToken(t *testing.T) {
	println(CreateToken("r3inbowari"))
}

func TestGetID(t *testing.T) {
	println(getIdByToken(CreateToken("r3inbowari")))
}

func TestAddAdmin(t *testing.T) {
	InitDB()
	db.Save(User{
		Uid:      "admin",
		Username: "admin",
		Password: utils.CreateMD5("15598870762"),
		Sex:      0,
		Role:     "admin",
		Avatar:   "qq.com",
		Status:   2,
	})
}
