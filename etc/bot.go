package etc

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var bot Bot

func GetBot() *Bot {
	return &bot
}

func (bot *Bot) ParseConfig(path string) {
	v, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	err = json.Unmarshal(v, bot)
	if err != nil {
		panic(err.Error())
	}
}

func (bot *Bot) Start() {

	//create bot
	dg, err := discordgo.New("Bot " + bot.Token)
	if err != nil {
		panic(err.Error())
	}

	//setup intends
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	//add handlers
	dg.AddHandler(bot.messageCreate)

	//establish connection to discord
	err = dg.Open()
	if err != nil {
		panic(err.Error())
	}

	//set status
	err = dg.UpdateGameStatus(0, bot.Status)
	if err != nil {
		panic(err.Error())
	}

	bot.run(dg)
}

func (bot *Bot) run(dg *discordgo.Session) {
	fmt.Printf("%s is now running, press CTRL-C to stop", bot.Name)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	err := dg.Close()
	if err != nil {
		return
	}
}

func (bot *Bot) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	//check if message is from itself
	if m.Author.ID == s.State.User.ID {
		return
	}
	args := strings.Split(m.Content, " ")
	for i, k := range Commands {
		if strings.ToLower(i.Name) == strings.ToLower(args[0][1:]) {
			k(s, m)
		}
	}

}
