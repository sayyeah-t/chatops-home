package slack

import (
	sl "github.com/nlopes/slack"
	"github.com/sayyeah-t/take2-chatops/src/config"
	"github.com/sayyeah-t/take2-chatops/src/opsdriver"
	"github.com/sayyeah-t/take2-chatops/src/util"
)

type SlackInterface struct {
	opsDriver opsdriver.DriverInterface
	confMap   map[string]string
	api       *sl.Client
	rtm       *sl.RTM
	runFlag   bool
}

func Init() *SlackInterface {
	si := &SlackInterface{}
	si.runFlag = true
	si.confMap = config.GetConfig()["slack"]
	si.auth()
	return si
}

func (si *SlackInterface) SetOpsDriver(driver opsdriver.DriverInterface) {
	si.opsDriver = driver
}

func (si *SlackInterface) Run() {
	go si.rtm.ManageConnection()
	for si.runFlag {
		select {
		case msg := <-si.rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *sl.MessageEvent:
				//println("Catched Message Event!")
				si.handleCommand(ev.Text)
			case *sl.InvalidAuthEvent:
				println("Invalid credentials")
				si.runFlag = false
			}
		}
	}
}

func (si *SlackInterface) Stop() {
	si.runFlag = false
}

func (si *SlackInterface) PostMessage(msg string) {
	si.rtm.SendMessage(si.rtm.NewOutgoingMessage(msg, si.confMap["channel"]))
}

func (si *SlackInterface) IsAvailable() error {
	_, err := si.api.GetUserInfo("U023BECGF")
	if err != nil {
		return err
	}
	return nil
}

func (si *SlackInterface) auth() {
	si.api = sl.New(si.confMap["bot_token"])
	si.rtm = si.api.NewRTM()
}

func (si *SlackInterface) handleCommand(command string) {
	si.PostMessage("Input: \"" + command + "\"")
	if util.IsCommand(command) {
		resp := si.opsDriver.DoCommand(util.ParseCommand(command))
		if resp != "" {
			si.PostMessage(resp)
		}
		si.PostMessage("Complete!")
	}
}
