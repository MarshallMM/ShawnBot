package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
	gl    []string
	vl    []string
)

const KuteGoAPIURL = "https://kutego-api-xxxxx-ew.a.run.app"

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}
func main() {
	//var gameList []string
	//var vetoList []string
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

type Gopher struct {
	Name string `json: "name"`
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// This function will be called (due to AddHandler above) every time a new
	// message is created on any channel that the authenticated bot has access to.
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	message := "hmm"
	var err error
	err = nil
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "gm" {
		message = "gm"
	}
	if m.Content == "DN" {
		message = "DEEZ NUTS"
	}
	if strings.Index(m.Content, "dn") != -1 {
		message = "DEEZ NUTS"
	}
	if strings.Index(m.Content, "DN") != -1 {
		message = "DEEZ NUTS"
	}
	if strings.Index(m.Content, "ligma") != -1 {
		message = "LIGMA BALLS"
	}
	if m.Content == "testTrout" {
		err = s.GuildMemberDelete(m.GuildID, m.Author.ID)
		message = "trout"
	}
	if strings.Index(m.Content, "true") != -1 {
		err = s.GuildMemberDelete(m.GuildID, m.Author.ID)
		message = "trout"
	}
	if strings.Index(m.Content, "where") != -1 {
		message = "I said where"
	}
	if m.Content == "totm" {
		message = "totm"
	}
	fmt.Println(m.Author)
	fmt.Println(m.GuildID)
	if message != "hmm" {
		// Send a text message
		_, err = s.ChannelMessageSend(m.ChannelID, message)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		return
	}
	return
}
