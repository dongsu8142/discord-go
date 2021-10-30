package commands

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/hands8142/discordhandler"
)

func Help(h *discordhandler.Handler, ds *discordgo.Session, dm *discordgo.Message, ctx *discordhandler.Context) {
	cp := h.Prefix
	maxlen := 0
	keys := make([]string, 0, len(h.Routes))
	cmdmap := make(map[string]*discordhandler.Route)

	for _, v := range h.Routes {
		if v.Description == "" {
			continue
		}
		l := len(v.Pattern)
		if l > maxlen {
			maxlen = l
		}

		cmdmap[v.Pattern] = v
		if v.Pattern == "help" || v.Pattern == "about" {
			continue
		}

		keys = append(keys, v.Pattern)
	}

	sort.Strings(keys)
	resp := "```autoit\n"

	v, ok := cmdmap["help"]
	if ok {
		keys = append([]string{v.Pattern}, keys...)
	}

	v, ok = cmdmap["about"]
	if ok {
		keys = append([]string{v.Pattern}, keys...)
	}

	for _, k := range keys {
		v := cmdmap[k]
		resp += fmt.Sprintf("%s%-"+strconv.Itoa(maxlen)+"s # %s\n", cp, v.Pattern+v.Help, v.Description)
	}

	resp += "```\n"

	ds.ChannelMessageSend(dm.ChannelID, resp)
}
