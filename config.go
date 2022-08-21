package main

import "github.com/tkanos/gonfig"

type Configuration struct {
	BotToken string
}

func GetConfiguration() Configuration {
	configuration := Configuration{}
	if err := gonfig.GetConf("config.json", &configuration); err != nil {
		panic(err)
	}
	if configuration.BotToken == "" {
		panic("Missing configuration 'BotToken'")
	}
	return configuration
}
