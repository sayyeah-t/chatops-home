package main

import (
	"github.com/sayyeah-t/chatops-home/src/chat"
	"github.com/sayyeah-t/chatops-home/src/config"
)

func main() {
	println("Hello, world!")
	if config.Init() != nil {
		println("load config error")
		return
	}
	//config.DumpConfig()
	if err := chat.Init(); err != nil {
		println(err.Error())
		return
	}
	chat.Run()
}
