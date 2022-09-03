package main

import (
	"flag"
	"fmt"
	"math/rand"
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
	tts := false
	var err error
	err = nil
	if m.Author.ID == s.State.User.ID {
		return
	}
	mes := strings.ToLower(m.Content)
	fmt.Println(mes)
	if m.Author.ID == "9382" {

		if rand.Intn(10) == 1 {
			message = "Ligma Balls"
			tts = true
		} else {
			return
		}
	}
	if strings.Index(mes, "gm") != -1 {
		message = "Gm"
	} else if strings.Index(mes, "gn") != -1 {
		message = "Gn"
	} else if strings.Index(mes, "gnn") != -1 {
		message = "Gnn"
	} else if strings.Index(mes, "totm") != -1 {
		message = "Totm"
	} else if strings.Index(mes, "trout") != -1 {
		message = "Trout that"
	} else if strings.Index(mes, "honk") != -1 {
		message = "HONK HONK"
	} else if strings.Index(mes, "beep") != -1 {
		message = "BEEP BEEP"
	}
	if strings.Index(mes, "dn") != -1 {
		message = "DEEZ NUTS"

	} else if strings.Index(mes, "ligma") != -1 {
		message = "LIGMA BALLS"

	}
	if m.Content == "testTrout" {
		err = s.GuildMemberDelete(m.GuildID, m.Author.ID)
		message = "trout"
		tts = true
	}
	if strings.Index(mes, "true") != -1 {
		err = s.GuildMemberDelete(m.GuildID, m.Author.ID)
		message = "trout"
	}
	if strings.Index(mes, "where") != -1 {
		message = "I said where"
	}

	fmt.Println(m.Author)

	if message != "hmm" {
		// Send a text message
		if rand.Intn(20) == 1 {
			tts = true
		}
		if tts {
			_, err = s.ChannelMessageSendTTS(m.ChannelID, message)
		} else {

			_, err = s.ChannelMessageSend(m.ChannelID, message)
		}

		if err != nil {
			fmt.Println(err)
		}
	} else {
		return
	}
	return
}
