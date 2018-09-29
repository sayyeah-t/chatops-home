package driver

import (
	"github.com/sayyeah-t/take2-chatops/src/config"
)

type Driver struct {
	confMap map[string]string
}

func Init() *Driver {
	d := &Driver{}
	d.confMap = config.GetConfig()["serverops"]
	return d
}

func (d *Driver) DoCommand(command []string) string {
	resp := ""
	switch command[0] {
	case "!health":
		resp = d.health(command)
	}
	return resp
}

func (d *Driver) health(command []string) string {
	if len(command) != 1 {
		if command[1] != d.confMap["nodename"] {
			return ""
		}
	}
	return d.confMap["nodename"] + "は元気みたいです〜！"
}
