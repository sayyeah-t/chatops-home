package main

import (
	"github.com/sayyeah-t/chatops-home/src/config"
)

func main() {
	println("Hello, world!")
	err := config.Init()
	if err != nil {
		println("load config error...")
		return
	}
	config.DumpConfig()
}
