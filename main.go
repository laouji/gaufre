package main

import (
	"./commands"
	"./config"
	"github.com/m0t0k1ch1/ape"
	"log"
	"net/http"
)

func main() {
	conf := config.LoadConfig()
	con := ape.NewConnection(conf.Nickname, conf.Username)
	con.UseTLS = true
	con.Password = conf.Password
	prefix := conf.MessagePrefix

	if err := con.Connect(conf.Server); err != nil {
		log.Fatal(err)
	}

	con.RegisterChannel(conf.Channel)

	con.AddAction("call-me", func(e *ape.Event) {
		con.Whois("laouji")
		message := conf.AdminUsername + "が今いないけど、IRCに呼んどいたよ"
		statusCode, err := commands.SendPush(e.Command().Args())

		if err != nil || statusCode != http.StatusOK {
			message = "ごめん！" + conf.AdminUsername + "をIRCに呼ぼうとしたら、こんなエラーが出た: " + err.Error()
		}

		con.SendMessage(prefix + message)
	})

	con.Loop()
}
