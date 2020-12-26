package core

import (
	"encoding/json"
	"momoEngine/utils"
	"testing"
)

func TestPage_Save(t *testing.T) {
	InitDB()

	var arr = []string{
		"721364",
		"464326",
	}
	fm, _ := json.Marshal(&arr)

	p := Page{
		Id:        utils.CreateUUID(),
		Name:      "Home",
		CookieIds: string(fm),
		Head: true,
	}

	p.Save()
}
