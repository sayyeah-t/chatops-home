package driver

import (
	"github.com/sayyeah-t/take2-chatops/src/config"
	"os/exec"
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
	case "!shutdown":
		resp = d.shutdown(command)
	case "!reboot":
		resp = d.reboot(command)
	}
	return resp
}

func (d *Driver) health(command []string) string {
	if len(command) == 2 {
		if command[1] != d.confMap["nodename"] {
			return ""
		}
	}
	return d.confMap["nodename"] + "は元気みたいです〜！"
}

func (d *Driver) shutdown(command []string) string {
	if len(command) == 2 {
		if command[1] == d.confMap["nodename"] {
			err := exec.Command("shutdown", "-h", "-t", "10").Run()
			if err != nil {
				return err.Error()
			}
			return d.confMap["nodename"] + "は10秒後にシャットダウンされます〜♪"
		}
	}
	return ""
}

func (d *Driver) reboot(command []string) string {
	if len(command) == 2 {
		if command[1] == d.confMap["nodename"] {
			err := exec.Command("shutdown", "-r", "-t", "10").Run()
			if err != nil {
				return err.Error()
			}
			return d.confMap["nodename"] + "は10秒後にリブートされます〜♪"
		}
	}
	return ""
}
