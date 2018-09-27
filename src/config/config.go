package config

import (
	"github.com/go-ini/ini"
)

var (
	configPath = "/etc/chatops-home/chatops-home.conf"
)

func Init() error {
	err := loadConf()
	if err != nil {
		return err
	}
	return nil
}

func GetConfig() map[string]map[string]string {
	return configurations
}

func DumpConfig() {
	for section, content := range configurations {
		println("[" + section + "]")
		for key, value := range content {
			println(key + " = " + value)
		}
	}
}

func loadConf() error {
	cfg, err := ini.InsensitiveLoad(configPath)
	if err != nil {
		return err
	}
	for name := range configurations {
		section, err := cfg.GetSection(name)
		if err != nil {
			return err
		}
		for key := range configurations[name] {
			data, err := section.GetKey(key)
			if err != nil {
				return err
			}
			configurations[name][key] = data.MustString("none")
		}
	}
	return nil
}
