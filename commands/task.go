package commands

import (
	"discord-bot/db"
	"discord-bot/utils"

	"github.com/bwmarrin/discordgo"
)

func HandleTask(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := utils.GetOptions(i.ApplicationCommandData().Options)

	description := options["descripcion"]

	task := db.Task{UserId: s.State.User.ID, Description: description.StringValue()}
	task.SaveTask()

	utils.SendToSource(s, i, "Tarea agregada con éxito")
}
