package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/hands8142/discordhandler"
)

func Ping(h *discordhandler.Handler, s *discordgo.Session, m *discordgo.Message, ctx *discordhandler.Context) {
	_, err := s.ChannelMessageSend(m.ChannelID, "Pong!")
	if err != nil {
		log.Print("err", err)
	}
}
