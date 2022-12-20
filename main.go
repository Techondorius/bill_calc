package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

var (
	Token     string
	BotPrefix = "!"
)

func ReadConfig() error {
	fmt.Println("Reading config file...")
	Token = os.Getenv("LINE_BOT_TOKEN")

	return nil

}

var BotId string

// var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	goBot.Identify.Intents = discordgo.IntentMessageContent

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)

	goBot.Identify.Intents = discordgo.IntentsGuildMessages

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running !")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}

	if m.Content == BotPrefix+"ping" {
		fmt.Println("sending")
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}
}

func main() {
	err := ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Start()

	<-make(chan struct{})
}
