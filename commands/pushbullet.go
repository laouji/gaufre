package commands

import (
	"../config"
	"../utils"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	apiBaseUrl = "https://api.pushbullet.com"
	headers    map[string]string
)

type PushReqParams struct {
	Email string `json:"email"`
	Type  string `json:"type"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func SendPush(args []string) (int, error) {
	conf := config.LoadConfig()

	headers = make(map[string]string)
	headers["Access-Token"] = conf.PushbulletAccessToken

	params := SetPushReqParams(conf, args)

	res, err := utils.HttpPostJson(apiBaseUrl+"/v2/pushes", headers, params)
	if err != nil {
		return 0, err
	}

	if res.StatusCode != http.StatusOK {
		rawBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return 0, err
		}
		log.Println(fmt.Printf("%s\n", rawBody))
		return res.StatusCode, fmt.Errorf(res.Status)
	}

	return res.StatusCode, nil
}

func SetPushReqParams(conf *config.ConfData, args []string) PushReqParams {
	msgBody := strings.Join(args[:], " ")
	if len(args) == 0 {
		msgBody = "(blank)"
	}

	return PushReqParams{
		conf.PushbulletUserEmail,
		"note",
		"You were mentioned in " + conf.Channel,
		msgBody,
	}
}
