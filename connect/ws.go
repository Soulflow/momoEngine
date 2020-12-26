package connect

import (
	"encoding/json"
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"golang.org/x/net/websocket"
	"log"
	"momoEngine/utils"
	"time"
)

var origin = "http://127.0.0.1:19980/"
var url = "ws://127.0.0.1:19980/ws"

type Message struct {
	ID   string      `json:"id"`
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Ts   time.Time   `json:"ts"`
	Data interface{} `json:"data"`
}

type UpdateInfo struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Patch int `json:"patch"`
}

var ws *websocket.Conn
var wsStatus = false

func InitWSC() {
	readyReconnect()
}

func readyReconnect() {
	var err error
	ws, err = websocket.Dial(url, "", origin)
	if err != nil {
		utils.Fail("[WS] connect failed")
		utils.Info("[WS] reconnect 5s later")
		time.Sleep(time.Second * 5)
		go readyReconnect()
		return
	}
	utils.Info("[WS] connected")
	wsStatus = true

	_ = auth()
	CheckUpdate()

	go hb()
}

func auth() error {
	if utils.GetConfig().Finger == "" {
		utils.Info("First Enter")
		ms := Message{
			Code: 407,
			Msg:  "getFinger",
			Ts:   time.Now(),
			Data: nil,
		}
		js, _ := json.Marshal(&ms)
		_, err := ws.Write(js)
		if err != nil {
			return err
		}

		var msg1 = make([]byte, 512)
		_, err = ws.Read(msg1)
		if err != nil {
			log.Fatal(err)
		}
		j, _ := jsonvalue.Unmarshal(msg1)
		v, _ := j.Get("data")
		utils.Info(v.String())
		utils.GetConfig().Finger = v.String()
		_ = utils.GetConfig().SetConfig()
	} else {
		utils.Info("finger " + utils.GetConfig().Finger)
	}
	return nil
}

func hb() {
	heartbeat := Message{
		ID:   utils.GetConfig().Finger,
		Code: 404,
		Msg:  "heartbeat",
		Ts:   time.Now(),
		Data: nil,
	}
	js, _ := json.Marshal(&heartbeat)

	for {
		_, err := ws.Write(js)
		if err != nil {
			utils.Warn("[WS] disconnected")
			wsStatus = false
			readyReconnect()
			return
		}
		time.Sleep(time.Second * 2)
	}
}

func CheckUpdate() {
	ms := Message{
		Code: 404,
		Msg:  "checkUpdate",
		Ts:   time.Now(),
		Data: nil,
	}
	js, _ := json.Marshal(&ms)
	_, err := ws.Write(js)
	if err != nil {
		log.Fatal(err)
	}

	var msg = make([]byte, 512)
	_, err = ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}

	j, _ := jsonvalue.Unmarshal(msg)
	major, _ := j.Get("data", "major")
	minor, _ := j.Get("data", "minor")
	patch, _ := j.Get("data", "patch")
	utils.Info("remote release version: " + major.String() + "." + minor.String() + "." + patch.String())
}
