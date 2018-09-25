package config

var (
	configurations = map[string]map[string]string{
		"default": {
			"nodename": "",
		},
		"slack": {
			"channel":   "",
			"bot_token": "",
		},
	}
)
