package commands

import (
	config "discord-bot/config"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	configuration config.Configuration

	dmPermissionFalse = false
	minCantidad       = 1.0

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
		{
			Name:         "pick",
			Description:  "Permite obtener cantidad específica de elementos aleatoreos",
			DMPermission: &dmPermissionFalse,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "lista",
					Description: "Obtiene cantidad específica aleatoreo de una lista de opciones separada por espacios",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "opciones",
							Description: "Opciones separadas por espacio",
							Required:    true,
						},
						{
							Type:        discordgo.ApplicationCommandOptionInteger,
							Name:        "cantidad",
							Description: "Cantidad de elementros a devolver, si no se especifíca se devuelve un único elemento",
							MinValue:    &minCantidad,
						},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "rol",
					Description: "Obtiene usuario aleatoreo los usuarios que pertenecen a un rol",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionRole,
							Name:        "rol",
							Description: "Rol a ordenar aleatoreamente",
							Required:    true,
						},
						{
							Type:        discordgo.ApplicationCommandOptionInteger,
							Name:        "cantidad",
							Description: "Cantidad de elementros a devolver, si no se especifíca se devuelve un único elemento",
							MinValue:    &minCantidad,
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
						{
							Type:        discordgo.ApplicationCommandOptionInteger,
							Name:        "cantidad",
							Description: "Cantidad de elementros a devolver, si no se especifíca se devuelve un único elemento",
							MinValue:    &minCantidad,
						},
					},
				},
			},
		},
		{
			Name:        "tarea",
			Description: "Permite realizar acciones sobre tareas",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "agregar",
					Description: "Permite registrar una tarea realizada",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "descripcion",
							Description: "Descripción de la tarea",
							Required:    true,
						},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "listado",
					Description: "Permite listar las tareas realizadas",
				},
			},
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"shuffle": HandleShuffle,
		"pick":    HandlePick,
		"tarea":   HandleTask,
	}
)

func init() {
	configuration = config.GetConfiguration()
}

func RegisterCommands(s *discordgo.Session) []*discordgo.ApplicationCommand {
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, command := range commands {
		registeredCommand, err := s.ApplicationCommandCreate(configuration.AppId, "", command)
		if err != nil {
			fmt.Printf("Cannot create '%s' command\n", command.Name)
			panic(err)
		}
		registeredCommands[i] = registeredCommand
	}
	return registeredCommands
}

func UnregisterCommands(s *discordgo.Session, commands []*discordgo.ApplicationCommand) {
	for _, registeredCommand := range commands {
		err := s.ApplicationCommandDelete(configuration.AppId, "", registeredCommand.ID)
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
