package driver

import (
	"github.com/sayyeah-t/take2-chatops/src/config"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

var (
	splatnet2statinkConfig = []string{
		"api_errors",
		"api_key",
		"cookie",
		"session_token",
		"user_lang",
	}
)

type Driver struct {
	confMap   map[string]string
	upload    bool
	latestLog string
}

func Init() *Driver {
	d := &Driver{}
	d.confMap = config.GetConfig()["splatoon"]
	d.upload = false
	d.latestLog = "Initialized..."
	os.Chdir(d.confMap["uploader_path"])
	d.initConfig()
	go d.uploadLoop()
	return d
}

func (d *Driver) DoCommand(command []string) string {
	resp := ""
	switch command[0] {
	case "!help":
		resp = d.helpMessage()
	case "!once":
		d.updateUploader()
		resp = d.uploadOnce(command)
	case "!start":
		d.updateUploader()
		resp = d.startUpload(command)
	case "!stop":
		resp = d.stopUpload(command)
	case "!status":
		resp = d.getStatus()
	}
	return resp
}

func (d *Driver) initConfig() {
	jsonText := "{\n"
	num := len(splatnet2statinkConfig) - 1
	for i, arg := range splatnet2statinkConfig {
		jsonText = jsonText + "\t\"" + arg + "\": \"\""
		if i != num {
			jsonText = jsonText + ","
		}
		jsonText = jsonText + "\n"
	}
	jsonText = jsonText + "}"
	ioutil.WriteFile("config.txt", []byte(jsonText), os.ModePerm)
	d.updateConfigArg("user_lang", "ja-JP")
	d.updateConfigArg("cookie", d.confMap["iksm_session"])
	d.updateConfigArg("api_key", d.confMap["api_key"])
}

func (d *Driver) updateToken(token string) error {
	tmpUpload := d.upload
	if d.upload {
		d.upload = false
	}
	err := d.updateConfigArg("cookie", token)
	if err != nil {
		return err
	}
	d.upload = tmpUpload
	return nil
}

func (d *Driver) updateConfigArg(key string, value string) error {
	err := exec.Command(
		"sed",
		"-i",
		"s/\""+key+"\": \".*\"/\""+key+"\": \""+value+"\"/g",
		d.confMap["uploader_path"]+"/config.txt",
	).Run()
	if err != nil {
		return err
	}
	return nil
}

func (d *Driver) startUpload(command []string) string {
	var msg string
	switch len(command) {
	case 1:
		msg = "戦績データをアップロードするぜ！"
		d.upload = true
	case 2:
		if err := d.updateToken(command[1]); err != nil {
			return err.Error()
		}
		msg = "iksm_sessionの更新完了！\n戦績データをアップロードするぜ！"
		d.upload = true
	default:
		msg = "コマンドの使い方間違ってんぞ！"
	}
	return msg
}

func (d *Driver) stopUpload(command []string) string {
	var msg string
	switch len(command) {
	case 1:
		if d.upload {
			msg = "戦績アップロード終了！"
			d.upload = false
		} else {
			msg = "今、アップロードはしてないっぽいぞ？"
		}
	default:
		msg = "コマンドの使い方間違ってんぞ！"
	}
	return msg
}

func (d *Driver) uploadLoop() {
	for {
		if d.upload {
			d.execUpload()
			time.Sleep(85 * time.Second)
		}
		time.Sleep(5 * time.Second)
	}
}

func (d *Driver) execUpload() {
	out, err := exec.Command(
		"python",
		"splatnet2statink.py",
		"-r",
	).Output()
	if err != nil {
		d.latestLog = err.Error()
	} else {
		d.latestLog = string(out)
	}
	println(string(out))
}

func (d *Driver) updateUploader() {
	err := exec.Command(
		"git",
		"pull",
		"origin",
		"master",
	).Run()
	if err != nil {
		println(err.Error())
		return
	}
	println("Update splatnet2statink completed!")
}

func (d *Driver) getStatus() string {
	if d.upload {
		return d.latestLog
	}
	return "今、アップロードはしてないっぽいぞ？"
}

func (d *Driver) helpMessage() string {
	msg := "コマンドはこんな感じで使ってくれよな！\n"
	msg = msg + "```\n"
	msg = msg + "!start [iksm_session]\n"
	msg = msg + "!stop\n"
	msg = msg + "!once [iksm_session]\n"
	msg = msg + "!status\n"
	msg = msg + "```"
	return msg
}

func (d *Driver) uploadOnce(command []string) string {
	if d.upload {
		return "アップロードモード中にこのコマンドは使えないぞー"
	}
	var msg string
	switch len(command) {
	case 1:
		msg = "新しいリザルトをアップロードするぜ！"
	case 2:
		if err := d.updateToken(command[1]); err != nil {
			return err.Error()
		}
		msg = "iksm_sessionの更新完了！\n新しいリザルトをアップロードするぜ！"
	default:
		msg = "コマンドの使い方間違ってんぞ！"
	}
	d.execUpload()
	return msg
}
