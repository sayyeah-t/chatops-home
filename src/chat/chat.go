package chat

import (
	"errors"
	"github.com/sayyeah-t/take2-chatops/src/chat/slack"
	"github.com/sayyeah-t/take2-chatops/src/config"
	"github.com/sayyeah-t/take2-chatops/src/opsdriver"
)

var (
	chatInterface ChatInterface
)

type ChatInterface interface {
	SetOpsDriver(opsdriver.DriverInterface)
	IsAvailable() error
	Run()
	Stop()
	PostMessage(string)
}

func Init(driver opsdriver.DriverInterface) error {
	conf := config.GetConfig()
	switch conf["default"]["chat_driver"] {
	case "slack":
		chatInterface = slack.Init()
	default:
		return errors.New("Specified driver not found")
	}
	chatInterface.SetOpsDriver(driver)
	return nil
}

func Run() {
	if chatInterface.IsAvailable() != nil {
		chatInterface.Run()
	}
}
