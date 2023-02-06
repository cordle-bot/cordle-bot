package main

import (
	"github.com/bwmarrin/discordgo"
	"cordle/config"
	"cordle/util"
	"time" // Temporary
)

// Path to read config from
const ConfigPath = "config/config.json"

func main() {
	// Load config file
	config := config.LoadConfig(ConfigPath)

	// Start discord bot
	discord, err := discordgo.New("Bot " + config.Token)
	util.CheckError(err, "Failed to initialise discord session")
	discord.Open()
	defer discord.Close()

	// Set the bot's status
	err = discord.UpdateGameStatus(0, config.Status)
	util.CheckError(err, "Failed to set status")

	// Temporary, stops the bot from instantly logging out
	time.Sleep(8 * time.Second)
}