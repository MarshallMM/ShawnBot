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
	"time"

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
func dojo() string {
	rand.Seed(time.Now().UnixNano())

	// Open the text file in the same directory as the compiled binary
	fileBytes, err := ioutil.ReadFile("./dojo.txt")
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
	} else if strings.Contains(mes, "dojo") {
		message = dojo()
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
	} else if strings.Contains(mes, "key") {
		message = `Once upon a time, there was a man named Shawn. He was a simple man with a simple life. He lived in a small town and worked at a local grocery store. Shawn had always been content with his life, but he often felt like there was something missing.

		One day, Shawn was walking home from work when he stumbled upon a strange object lying on the ground. It was a small, silver key. Shawn picked it up and examined it closely. He had no idea where it came from or what it unlocked, but he felt like it was important.
		
		Over the next few days, Shawn couldn't stop thinking about the key. He decided to take a day off from work and try to find out where it belonged. He walked all over town, trying the key in every lock he could find. But none of them worked.
		
		Just when Shawn was about to give up, he stumbled upon an old, abandoned house at the edge of town. He had never seen it before, and it looked like it hadn't been inhabited for years. There was an old lock on the front door, and without even thinking, Shawn inserted the key.
		
		To his surprise, it worked! The lock clicked open, and Shawn cautiously stepped inside. The house was dusty and dimly lit, but it had an eerie sense of familiarity. Shawn walked through the rooms, examining the old furniture and decorations. But there was one room in particular that caught his attention.
		
		It was a small room at the back of the house. The walls were lined with shelves, and there were countless books and artifacts scattered throughout the room. Shawn felt drawn to a large, leather-bound book sitting on a pedestal in the center of the room. He walked up to it and began to read.
		
		As he flipped through the pages, Shawn realized that the book was a journal. But it wasn't just any journal - it was the journal of the previous owner of the house. The owner had been a traveler and adventurer, and the journal was filled with stories of his exploits and discoveries.
		
		As Shawn read through the journal, he became more and more engrossed in the stories. He felt like he was living the adventures himself, experiencing the thrill of discovery and the danger of unknown lands. And then he came across a particularly intriguing passage.
		
		It was a map - a map of a hidden city deep in the jungle. According to the journal, the city was filled with treasures and artifacts beyond imagination. Shawn felt a rush of excitement and determination. He knew he had to find this city.
		
		Without hesitation, Shawn left the old house and began his journey. He traveled through dense forests, across raging rivers, and up treacherous mountains. He faced dangerous animals and harsh weather, but he never gave up. He was driven by the promise of adventure and discovery.
		
		Finally, after weeks of travel, Shawn saw it - the hidden city. It was more magnificent than he ever could have imagined. He walked through the streets, marveling at the ancient buildings and artifacts. He collected treasures and examined relics, feeling like he was living a dream.
		
		But as Shawn was about to leave the city, he heard a faint whisper. It was a voice he recognized - the voice of the previous owner of the house. The voice told him that the city was not meant for outsiders, and that he must leave immediately.
		
		Shawn was torn. He didn't want to leave the city, but he also didn't want to disrespect the wishes of the previous owner. In the end, he decided to leave, but not before taking one last relic - a small, silver key.
		
		As Shawn made his way back home, he couldn't stop thinking about the city and the adventures he had experienced. But he also couldn't shake the feeling that he had done something wrong by taking the key and the relic.
		
		When he arrived back in town, Shawn went straight to the old house and placed the key and the relic back where he had found them. He felt a sense of relief and closure, knowing that he had done the right thing.
		
		From that day forward, Shawn felt a renewed sense of purpose in his life. He realized that he didn't need to travel to far-off lands to experience adventure and excitement - he could find it in his own town and in his everyday life. And he was content, knowing that he had lived an adventure that he would never forget, but also that he had done the right thing in the end.`
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
