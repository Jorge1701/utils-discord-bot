package utils

import "github.com/bwmarrin/discordgo"

func GetSubcommand(i *discordgo.InteractionCreate) string {
	options := i.ApplicationCommandData().Options
	if len(options) != 1 {
		return ""
	} else {
		return options[0].Name
	}
}

func GetOptions(options []*discordgo.ApplicationCommandInteractionDataOption) map[string]*discordgo.ApplicationCommandInteractionDataOption {
	optionsMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, option := range options {
		optionsMap[option.Name] = option
	}
	return optionsMap
}

func GetSubcommandOptions(i *discordgo.InteractionCreate) map[string]*discordgo.ApplicationCommandInteractionDataOption {
	options := i.ApplicationCommandData().Options[0].Options
	return GetOptions(options)
}
