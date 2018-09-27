package chat

import (
	"errors"
	"github.com/sayyeah-t/chatops-home/src/chat/slack"
	"github.com/sayyeah-t/chatops-home/src/config"
)

var (
	chatInterface ChatInterface
)

type ChatInterface interface {
	IsAvailable() error
	Run()
	Stop()
	PostMessage(string)
}

func Init() error {
	conf := config.GetConfig()
	switch conf["default"]["chat_driver"] {
	case "slack":
		chatInterface = slack.Init()
	default:
		return errors.New("Specified driver not found")
	}
	return nil
}

func Run() {
	if chatInterface.IsAvailable() != nil {
		chatInterface.Run()
	}
}
