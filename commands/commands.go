package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	dmPermissionFalse = false

	commands = []*discordgo.ApplicationCommand{
		{
			Name:         "shuffle",
			Description:  "Permite obtener listas aleatoreas de distintas maneras",
			DMPermission: &dmPermissionFalse,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "lista",
					Description: "Ordena aleatoreamente una lista de opciones separada por espacios",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "opciones",
							Description: "Opciones separadas por espacio",
							Required:    true,
						},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "rol",
					Description: "Ordena aleatoreamente los usuarios que pertenecen a un rol",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionRole,
							Name:        "rol",
							Description: "Rol a ordenar aleatoreamente",
							Required:    true,
						},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "audio",
					Description: "Ordena aleatoreamente lo usuarios conectados a un canal",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionChannel,
							Name:        "canal",
							Description: "Canal a ordenar aleatoreamente",
							Required:    true,
							ChannelTypes: []discordgo.ChannelType{
								discordgo.ChannelTypeGuildVoice,
							},
						},
					},
				},
			},
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"shuffle": HandleShuffle,
	}
)

func RegisterCommands(discord *discordgo.Session) []*discordgo.ApplicationCommand {
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, command := range commands {
		registeredCommand, err := discord.ApplicationCommandCreate(discord.State.User.ID, "", command)
		if err != nil {
			fmt.Printf("Cannot create '%s' command\n", command.Name)
			panic(err)
		}
		registeredCommands[i] = registeredCommand
	}
	return registeredCommands
}

func UnregisterCommands(discord *discordgo.Session, commands []*discordgo.ApplicationCommand) {
	for _, registeredCommand := range commands {
		err := discord.ApplicationCommandDelete(discord.State.User.ID, "", registeredCommand.ID)
		if err != nil {
			fmt.Printf("Cannot delete '%s' command\n", registeredCommand.Name)
		}
	}
}

func HandleCommands(s *discordgo.Session, i *discordgo.InteractionCreate) {
	handler, handlerExists := commandHandlers[i.ApplicationCommandData().Name]
	if handlerExists {
		handler(s, i)
	} else {
		fmt.Printf("There is no handler for the interaction: '%s'.\n", i.ApplicationCommandData().Name)
	}
}
