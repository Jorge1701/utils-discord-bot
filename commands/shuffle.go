package commands

import (
	utils "discord-bot/utils"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HandleShuffle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	subcommand := utils.GetSubcommand(i)
	options := utils.GetSubcommandOptions(i)

	var respuesta strings.Builder
	var lista []string

	switch subcommand {
	case "lista":
		lista = strings.Fields(options["opciones"].StringValue())
		respuesta.WriteString(fmt.Sprintf("Shuffle lista `%s`\n\n", strings.Join(lista, "` `")))
	case "rol":
		rol := options["rol"].RoleValue(s, i.GuildID)
		respuesta.WriteString(fmt.Sprintf("Shuffle rol %s\n\n", rol.Mention()))

		members := utils.GetMembersIdWithRole(s, i.GuildID, rol.ID)
		for _, member := range members {
			lista = append(lista, fmt.Sprintf("<@%s>", member))
		}
	case "audio":
		respuesta.WriteString("Shuffle audio aún no implementado :sweat_smile:")
	}

	if len(lista) > 0 {
		utils.ShuffleSlice(lista)

		for i, opcion := range lista {
			respuesta.WriteString(fmt.Sprintf("%d. %s\n", i+1, opcion))
		}
	} else {
		respuesta.WriteString("No se encontraron opciones :thinking:")
	}

	if respuesta.Len() > 0 {
		utils.SendToSource(s, i, respuesta.String())
	}
}
