package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	commands "discord-bot/commands"

	"github.com/bwmarrin/discordgo"
)

func main() {
	config := GetConfiguration()

	// Create session with BotToken
	discord, err := discordgo.New(fmt.Sprintf("Bot %s", config.BotToken))

	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		return
	}

	// Opens websocket and begins listening
	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}

	// Defer the closing of websocket connection, in case o panic
	defer discord.Close()

	// Identify the intents of your bot
	discord.Identify.Intents = discordgo.IntentsGuildMembers

	// Register handlers for user interactions
	discord.AddHandler(commands.HandleCommands)

	// Register commands for user interactions
	registeredCommands := commands.RegisterCommands(discord)

	// Initialization finished
	fmt.Println("Bot is now running.")

	// Wait for closing signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Remove commands
	commands.UnregisterCommands(discord, registeredCommands)
}
