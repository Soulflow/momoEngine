package core

import (
	"fmt"
	"testing"
)

func TestCookie_Save(t *testing.T) {
	InitDB()
	s := Cookie{
		Id:    "96648586-14ff-4749-97b6-e34b2d12b622",
		Name:  "Home",
		Key:   "token",
		Value: "1008",
		Type:  0,
	}
	s.Save()
}

func TestCookie_Update(t *testing.T) {
	InitDB()
	u := Cookie{
		Id:    "96648586-14ff-4749-97b6-e34b2d12b622",
		Value: "2003",
	}
	u.Update()
}

func TestCookie_Delete(t *testing.T) {
	InitDB()
	d := Cookie{
		Id: "96648586-14ff-4749-97b6-e34b2d12b622",
	}
	d.Delete()
}

func TestCookie_useragent(t *testing.T) {

	InitDB()
	d := Cookie{
		Id: "96648586-14ff-4749-97b6-e34b2d12b622",
	}
	d.Delete()

	s := Cookie{
		Id:    "96648586-14ff-4749-97b6-e34b2d12b622",
		Name:  "用户代理",
		Key:   "user-agent",
		Value: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36",
		Type:  CookieTypeDefault,
	}
	s.Save()
}

func TestCookiePreload(t *testing.T) {
	InitDB()
	CookiesPreload()

	s := Cookie{
		Id:    "96648586-14ff-4749-97b6-e34b2d12b611",
		Name:  "用户代理",
		Key:   "user-agent1",
		Value: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36",
		Type:  CookieTypeDefault,
	}
	s.CookiesAdd()

	s = Cookie{
		Id:    "96648586-14ff-4749-97b6-e34b2d12b612",
		Name:  "用户代理",
		Key:   "user-agent2",
		Value: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36",
		Type:  CookieTypeDefault,
	}
	s.CookiesAdd()

	s = Cookie{
		Id:    "96648586-14ff-4749-97b6-e34b2d12b613",
		Name:  "用户代理",
		Key:   "user-agent3",
		Value: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36",
		Type:  CookieTypeDefault,
	}
	s.CookiesAdd()
	fmt.Println(CookiesCount())
}

func TestCookiesCount(t *testing.T) {
	fmt.Println(CookiesCount())
}
