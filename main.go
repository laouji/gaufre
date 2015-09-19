package main

import (
	"./commands"
	"./config"
	"github.com/thoj/go-ircevent"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type Stash struct {
	Mentions []string
}

var stash = &Stash{
	Mentions: nil,
}

func main() {
	conf := config.LoadConfig()
	con := irc.IRC(conf.Nickname, conf.Username)
	con.UseTLS = true
	con.Password = conf.Password
	prefix := conf.MessagePrefix

	if err := con.Connect(conf.Server); err != nil {
		log.Fatal(err)
	}

	// RPL_WELCOME
	con.AddCallback("001", func(e *irc.Event) {
		con.Join(conf.Channel)
	})

	con.AddCallback("PRIVMSG", func(e *irc.Event) {
		message := e.Message()
		if targetName(message) != conf.AdminUsername {
			return
		}

		stash.Mentions = append(stash.Mentions, message)
		con.Whois(conf.AdminUsername)
	})

	// 301 - RPL_AWAY
	con.AddCallback("301", func(e *irc.Event) {
		botReply := conf.AdminUsername + "が今いないけど、メッセージ転送しといたよ"

		for mentionCount() > 0 {
			args := strings.Split(nextMention(), " ")
			statusCode, err := commands.SendPush(args)

			if err != nil || statusCode != http.StatusOK {
				botReply = "ごめん！" + conf.AdminUsername + "をIRCに呼ぼうとしたら、こんなエラーが出た: " + err.Error()
			}

			con.Privmsg(conf.Channel, prefix+botReply)
		}
	})

	// 318 - RPL_ENDOFWHOIS
	con.AddCallback("318", func(e *irc.Event) {
		stash.Mentions = stash.Mentions[:0]
	})

	con.Loop()
}

func targetName(msg string) string {
	pattern := `^([^:]+): `
	matches := regexp.MustCompile(pattern).FindStringSubmatch(msg)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func mentionCount() int {
	return len(stash.Mentions)
}

func nextMention() string {
	var mention = ""
	mention, stash.Mentions = stash.Mentions[0], stash.Mentions[mentionCount():]

	return mention
}
