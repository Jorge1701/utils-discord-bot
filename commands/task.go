package commands

import (
	"discord-bot/db"
	"discord-bot/utils"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HandleTask(s *discordgo.Session, i *discordgo.InteractionCreate) {
	subcommand := utils.GetSubcommand(i)
	options := utils.GetSubcommandOptions(i)

	switch subcommand {
	case "agregar":
		description := options["descripcion"]

		task := db.Task{UserId: s.State.User.ID, Description: description.StringValue()}
		task.SaveTask()

		utils.SendToSourceEphemeral(s, i, fmt.Sprintf("Tarea agregada `%s`", description.StringValue()))
	case "listado":
		tasks := db.ListTasks(s.State.User.ID)

		var response strings.Builder

		if len(tasks) == 0 {
			response.WriteString("No hay tareas registradas")
		} else {
			response.WriteString("Tareas:\n")
			for _, task := range tasks {
				response.WriteString(fmt.Sprintf("\n%s", task.Description))
			}
		}

		utils.SendToSourceEphemeral(s, i, response.String())
	}
}
