package main

import (
	"github.com/sayyeah-t/take2-chatops/src/app/serverops/driver"
	"github.com/sayyeah-t/take2-chatops/src/chat"
	"github.com/sayyeah-t/take2-chatops/src/config"
)

var (
	configPath      = "/etc/take2-chatops/serverops.conf"
	serveropsConfig = map[string]map[string]string{
		"serverops": {
			"nodename": "",
		},
	}
)

func main() {
	if config.InitWithAdditionalArgs(configPath, serveropsConfig) != nil {
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
