package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/hands8142/discordhandler"
	"hands8142.com/discord-go/commands"
	"hands8142.com/discord-go/config"
)

var Session, _ = discordgo.New()

func main() {
	const fileName = "./config.json"

	cfg, err := config.ParseConfigFromJSONFile(fileName)

	if err != nil {
		log.Fatalf("error config file parsing, %s\n", err)
	}

	Session.Token = "Bot " + cfg.Token

	Session.Identify.Intents = discordgo.IntentsAll

	registerCommands(Session, cfg)

	err = Session.Open()
	if err != nil {
		log.Fatalf("error opening connection to Discord, %s\n", err)
	}

	log.Printf("Now running. Press CTRL-C to exit...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	Session.Close()
}

func registerCommands(s *discordgo.Session, cfg *config.Config) {
	Router := discordhandler.New(cfg.Prefix)
	s.AddHandler(Router.OnMessageCreate)

	Router.Route("ping", "Pong", commands.Ping)
	Router.Route("help", "Display this message.", commands.Help)
}
