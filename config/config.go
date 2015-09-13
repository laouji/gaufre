package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ConfData struct {
	Username              string `yaml:"username"`
	Nickname              string `yaml:"nickname"`
	Password              string `yaml:"password"`
	Server                string `yaml:"server"`
	Channel               string `yaml:"channel"`
	AdminUsername         string `yaml:"admin_username"`
	MessagePrefix         string `yaml:"message_prefix"`
	PushbulletAccessToken string `yaml:"pushbullet_access_token"`
	PushbulletUserEmail   string `yaml:"pushbullet_user_email"`
}

func LoadConfig() *ConfData {
	buf, err := ioutil.ReadFile("config/config_local.yml")

	d := ConfData{}
	if err != nil {
		buf, err := ioutil.ReadFile("config/config.yml")
		if err != nil {
			panic(err)
		}

		if err := yaml.Unmarshal(buf, &d); err != nil {
			panic(err)
		}
		return &d
	}

	if err := yaml.Unmarshal(buf, &d); err != nil {
		panic(err)
	}

	return &d
}
