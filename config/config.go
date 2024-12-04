package config

import "github.com/tkanos/gonfig"

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST string
	DB_PORT string
	DB_NAME string
}// pembuatan struct untuk mendapatkan informasi database dari json

func GetConfig() Configuration {
	conf := Configuration{}
	gonfig.GetConf("config/config.json", &conf) //menggunakan json untuk config database
	return conf
}