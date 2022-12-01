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

	if m.Author.ID == "9382" {
		ranShawn := rand.Intn(100)
		if ranShawn == 1 {
			message = "Ligma Balls"
			tts = true
		} else if ranShawn == 2 {
			message = "Dude Shawn you are so fake"
			tts = true
		} else {
			return
		}
	}
	if strings.Contains(mes, "gm") {
		message = "Gm"
	} else if strings.Contains(mes, "gn") {
		message = "Gn"
	} else if strings.Contains(mes, "gnn") {
		message = "Gnn"
	} else if strings.Contains(mes, "totm") {
		message = "Totm"
	} else if strings.Contains(mes, "trout") {
		message = "Trout that"
		tts = true
	} else if strings.Contains(mes, "shit") {
		message = "Thats a Justin"
		tts = true
	} else if strings.Contains(mes, "shawn") {
		message = "If shawn has a million fans, then I am one of them. If shawn has ten fans, then I am one of them. If shawn has only one fan then that is me. If shawn has no fans, then that means I am no longer on earth. If the world is against shawn, then I am against the world."
	} else if strings.Contains(mes, "who") {
		message = `WHO is Shawn Whitmore?

		In geography, My World
		
		In reality, My Life
		
		In history, My King
		
		In mathematics, My Solution
		
		In mythology, My god
		
		In astronomy, My Universe
		
		If I'm Blind, He's Light
		
		If I'm Hungry, He's Food
		
		If I'm sick, He's Medicine
		
		For Me, He's Everything`
	} else if strings.Contains(mes, "honk") {
		message = "HONK HONK"
	} else if strings.Contains(mes, "beep") {
		message = "BEEP BEEP"
	} else if strings.Contains(mes, "cs") {
		message = "More like... cs NO!"
	}
	if strings.Contains(mes, "dn") {
		message = "DEEZ NUTS"

	} else if strings.Contains(mes, "ligma") {
		message = "LIGMA BALLS"

	} else if strings.Contains(mes, "sugondes") {
		message = "SUCK ON THESE NUTS"

	}
	if m.Content == "testTrout" {
		_ = s.GuildMemberDelete(m.GuildID, m.Author.ID)
		message = "trout"

	}
	if strings.Contains(mes, "true") {
		message = m.Author.Username + " said the wrong trout. Ban his ass."
		tts = true
	}
	if strings.Contains(mes, "where") {
		message = "I said where"
	}

	fmt.Println(m.Author, m.Content)

	if message != "hmm" {
		// Send a text message
		if rand.Intn(25) == 1 {
			tts = true
		}
		if tts {
			_, err = s.ChannelMessageSendTTS(m.ChannelID, message)
			fmt.Println("tts " + message)
		} else {

			_, err = s.ChannelMessageSend(m.ChannelID, message)
			fmt.Println(message)
		}

		if err != nil {
			fmt.Println(err)
		}
	} else {
		return
	}
}
