package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"io/ioutil"
	"log"
	"path/filepath"
	"time"

	"github.com/bwmarrin/discordgo"
)

// constants
const triggerDir string = "./src/triggers/"

// Variables used for command line parameters
var (
	Token           string
	cooldownEndTime time.Time
	triggers        []string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}
func ttsCooldownCheck(input bool) bool {
	cooldownDuration := 5 * time.Minute

	if time.Now().Before(cooldownEndTime) {
		return false
	}

	if input {
		// start the cooldown timer
		cooldownEndTime = time.Now().Add(cooldownDuration)
	}

	return input
}
func iterateTextFileResponces() []string {
	var filenames []string
	files, err := ioutil.ReadDir(triggerDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".txt" {
			name := filepath.Base(file.Name())
			name = name[:len(name)-4] // Remove the last 4 characters (i.e. ".txt")
			filenames = append(filenames, name)
		}
	}
	return filenames
}

func triggerResponse(trigger string) string {
	rand.Seed(time.Now().UnixNano())

	// Open the text file in the same directory as the compiled binary
	filepath := triggerDir + trigger + ".txt"
	fileBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	fileContent := string(fileBytes)

	// Split the file content into lines
	lines := strings.Split(fileContent, "\n")

	// Remove leading/trailing spaces from each line
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}

	// Choose a random line from the file
	randomIndex := rand.Intn(len(lines))
	randomLine := lines[randomIndex]

	// Print the randomly selected line to the console
	fmt.Println(randomLine)
	return randomLine
}

func main() {
	//Load triggers
	triggers = iterateTextFileResponces()
	fmt.Println(triggers)

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
		ranShawn := rand.Intn(20)
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

	//Highest Level of overriding
	if strings.Contains(mes, "true") {
		message = m.Author.Username + " said the wrong trout. Ban his ass."
		tts = true
	}

	fmt.Println(m.Author, m.Content)
	//Textfile triggers
	//triggers := iterateTextFileResponces()
	fmt.Println(triggers)
	for _, trigger := range triggers {
		if strings.Contains(mes, trigger) {
			newmessage := triggerResponse(trigger)
			_, _ = s.ChannelMessageSend(m.ChannelID, newmessage)
			fmt.Println(newmessage)
			time.Sleep(500 * time.Millisecond)
		}
	}
	if message != "hmm" {
		// Send a text message
		if rand.Intn(15) == 1 {
			tts = true
		}

		if tts && ttsCooldownCheck(tts) {
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
