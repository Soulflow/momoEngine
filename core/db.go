package core

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"momoEngine/utils"
)

var db *gorm.DB

func initTable() {
	utils.Info("[DB] tb init")
	if db.HasTable(&Cookie{}) {
		db.AutoMigrate(&Cookie{})
	} else {
		db.CreateTable(&Cookie{})
	}

	if db.HasTable(&Page{}) {
		db.AutoMigrate(&Page{})
	} else {
		db.CreateTable(&Page{})
	}

	if db.HasTable(&User{}) {
		db.AutoMigrate(&User{})
	} else {
		db.CreateTable(&User{})
	}

}

func InitDB() {
	utils.Info("[DB] db init")
	var err error
	db, err = gorm.Open("mysql", "root:159463@(r3inbowari.top:3306)/momo?charset=utf8&parseTime=True")
	if err != nil {
		utils.Info("[DB] init failed" + err.Error())
	}
	initTable()
}

func GetDB() *gorm.DB {
	return db
}
