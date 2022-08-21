package main

import (
	commands "discord-bot/commands"
	config "discord-bot/config"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	configuration := config.GetConfiguration()

	// Create discord session with BotToken
	ds, err := discordgo.New(fmt.Sprintf("Bot %s", configuration.BotToken))

	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		return
	}

	// Opens websocket and begins listening
	err = ds.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}

	// Defer the closing of websocket connection, in case of panic
	defer ds.Close()

	// Identify the intents of your bot
	ds.Identify.Intents = discordgo.IntentsGuildMembers

	// Register handlers for user interactions
	ds.AddHandler(commands.HandleCommands)

	// Register commands for user interactions
	registeredCommands := commands.RegisterCommands(ds)

	// Initialization finished
	fmt.Println("Bot is now running.")

	// Wait for closing signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Remove commands
	commands.UnregisterCommands(ds, registeredCommands)

	fmt.Println("Gracefully shutting down.")
}
