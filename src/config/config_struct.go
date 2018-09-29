package config

var (
	configurations = map[string]map[string]string{
		"default": {
			"chat_driver": "",
		},
		"slack": {
			"channel":   "",
			"bot_token": "",
		},
	}
)
