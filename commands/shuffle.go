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

		if guild, err := s.State.Guild(i.GuildID); err != nil {
			fmt.Printf("%s\n", err)
			respuesta.WriteString("Ocurrió un error :pensive:")
		} else {
			for _, member := range guild.Members {
				if utils.SliceContains(member.Roles, rol.ID) {
					if !utils.SliceContains(lista, member.Mention()) {
						lista = append(lista, member.Mention())
					}
				}
			}
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
