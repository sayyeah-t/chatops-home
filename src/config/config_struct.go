package config

var (
	configurations = map[string]map[string]string{
		"default": {
			"nodename":    "",
			"chat_driver": "",
		},
		"slack": {
			"channel":   "",
			"bot_token": "",
		},
	}
)
