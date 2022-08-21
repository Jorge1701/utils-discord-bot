package commands

import (
	"discord-bot/utils"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HandlePick(s *discordgo.Session, i *discordgo.InteractionCreate) {
	subcommand := utils.GetSubcommand(i)
	options := utils.GetSubcommandOptions(i)

	cantidad := int64(1)

	if cantidadSeleccionada, ok := options["cantidad"]; ok {
		cantidad = cantidadSeleccionada.IntValue()
	}

	var respuesta strings.Builder
	var lista []string

	switch subcommand {
	case "lista":
		lista = strings.Fields(options["opciones"].StringValue())
		respuesta.WriteString(fmt.Sprintf("Pick `%d` elementos de lista `%s`\n\n", cantidad, strings.Join(lista, "` `")))
	case "rol":
		rol := options["rol"].RoleValue(s, i.GuildID)
		respuesta.WriteString(fmt.Sprintf("Pick `%d` usuarios con rol %s\n\n", cantidad, rol.Mention()))

		members := utils.GetMembersIdWithRole(s, i.GuildID, rol.ID)
		for _, member := range members {
			lista = append(lista, fmt.Sprintf("<@%s>", member))
		}
	case "audio":
		respuesta.WriteString("Pick audio aún no implementado :sweat_smile:")
	}

	if len(lista) > 0 {
		utils.ShuffleSlice(lista)

		if cantidad == 1 {
			respuesta.WriteString(fmt.Sprintf("%s", lista[0]))
		} else {
			for i, opcion := range lista {
				if int64(i) >= cantidad {
					break
				}

				respuesta.WriteString(fmt.Sprintf("%d. %s\n", i+1, opcion))
			}
		}
	} else {
		respuesta.WriteString("No se encontraron opciones :thinking:")
	}

	if respuesta.Len() > 0 {
		utils.SendToSource(s, i, respuesta.String())
	}
}
