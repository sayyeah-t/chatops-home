package main

import (
	"github.com/sayyeah-t/take2-chatops/src/app/splatoon/driver"
	"github.com/sayyeah-t/take2-chatops/src/chat"
	"github.com/sayyeah-t/take2-chatops/src/config"
)

var (
	configPath     = "/etc/take2-chatops/splatoon.conf"
	splatoonConfig = map[string]map[string]string{
		"splatoon": {
			"iksm_session":  "",
			"api_key":       "",
			"uploader_path": "",
		},
	}
)

func main() {
	if config.InitWithAdditionalArgs(configPath, splatoonConfig) != nil {
		println("load config error")
		return
	}
	//config.DumpConfig()
	if err := chat.Init(driver.Init()); err != nil {
		println(err.Error())
		return
	}
	chat.Run()
}
