package momoEngine

import (
	"fmt"
	"momoEngine/connect"
	"momoEngine/core"
	"momoEngine/utils"
)

func InitSys() {
	Blue("                                  ______             _            ")
	Blue("                                 |  ____|           (_)           ")
	Blue("  _ __ ___   ___  _ __ ___   ___ | |__   _ __   __ _ _ _ __   ___ ")
	Blue(" | '_ ` _ \\ / _ \\| '_ ` _ \\ / _ \\|  __| | '_ \\ / _` | | '_ \\ / _ \\")
	Blue(" | | | | | | (_) | | | | | | (_) | |____| | | | (_| | | | | |  __/")
	Blue(" |_| |_| |_|\\___/|_| |_| |_|\\___/|______|_| |_|\\__, |_|_| |_|\\___|")
	Blue("                                                __/ |             ")
	Blue("                                               |___/              ")

	Blue(" -- --- -- --- . -. --. .. -. . version: 1.0.0 author: r3inbowari")
	Blue("")
	utils.Info("system init")
	utils.InitConfig()
	connect.InitWSC()
	core.InitDB()
	core.CookiesPreload()
	core.RunServer()
}

func Blue(msg string) {
	fmt.Printf("\x1b[%dm"+msg+" \x1b[0m\n", 34)
}
