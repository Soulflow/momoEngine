package core

import (
	"momoEngine/utils"
	"sync"
)

const (
	CookieTypeDefault  = 0
	CookieTypeConsumer = 1
)

type Cookie struct {
	Id    string `json:"id" gorm:"primary_key"`
	Type  int    `json:"type"`
	Name  string `json:"name"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (ck *Cookie) Save() {
	db.Create(ck)
}

func (ck *Cookie) Update() {
	db.Model(ck).Update(ck)

}
func (ck *Cookie) Delete() {
	db.Delete(ck)
}

var Cookies sync.Map

func CookiesPreload() {
	var cookies []Cookie
	db.Where("type = ?", CookieTypeDefault).Find(&cookies)
	for _, v := range cookies {
		Cookies.Store(v.Id, v)
		utils.Info("[Cookie] load default cookie -> " + v.Key)
	}

	db.Where("type = ?", CookieTypeConsumer).Find(&cookies)
	for _, v := range cookies {
		Cookies.Store(v.Id, v)
		utils.Info("[Cookie] load consumer cookie -> " + v.Key)
	}
}

func CookiesCount() int {
	var retCount = 0
	Cookies.Range(func(key, value interface{}) bool {
		retCount++
		return true
	})
	return retCount
}

func (ck *Cookie) CookiesAdd() {
	utils.Info("[Cookie] cookie add -> " + ck.Key)
	Cookies.Store(ck.Id, ck)
}

func CookiesClear() {
	utils.Warn("[Cookie] cookies gc")
	Cookies.Range(func(key, value interface{}) bool {
		Cookies.Delete(key)
		return true
	})
}

func CookiesList() []Cookie {
	ret := make([]Cookie, 0)
	Cookies.Range(func(key, value interface{}) bool {
		ret = append(ret, value.(Cookie))
		return true
	})
	return ret
}
