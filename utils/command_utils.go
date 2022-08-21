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

func GetSubcommandOptions(i *discordgo.InteractionCreate) map[string]*discordgo.ApplicationCommandInteractionDataOption {
	options := i.ApplicationCommandData().Options[0].Options
	optionsMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, option := range options {
		optionsMap[option.Name] = option
	}
	return optionsMap
}
