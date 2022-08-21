package config

import "github.com/tkanos/gonfig"

type Configuration struct {
	BotToken string
	AppId    string
}

func GetConfiguration() Configuration {
	configuration := Configuration{}
	if err := gonfig.GetConf("config.json", &configuration); err != nil {
		panic(err)
	}
	if configuration.BotToken == "" {
		panic("Missing configuration 'BotToken'")
	}
	if configuration.AppId == "" {
		panic("Missing configuration 'AppId'")
	}
	return configuration
}
