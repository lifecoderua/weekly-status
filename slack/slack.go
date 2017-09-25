package slack

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

const api = "https://slack.com/api/"

var tokenCache string

// Message recipient
func Message(channel string, message string) {
	log.Print("MSG<><><>")
	run("chat.postMessage", channel, message)
}

// func token() string {
// 	return os.Getenv("BOT_TOKEN")
// }

func run(endpoint string, channel string, text string) {
	uri := api + endpoint

	resp, err := http.PostForm(uri,
		url.Values{
			"token":   {token()},
			"channel": {channel},
			"text":    {text},
			"as_user": {"true"}})

	if nil != err {
		fmt.Println("errorination happened reading the body", err)
		return
	}

	defer resp.Body.Close()
}

func token() string {
	if tokenCache != "" {
		return tokenCache
	}

	tokenCache = os.Getenv("BOT_TOKEN")
	if tokenCache == "" {
		log.Panic("$BOT_TOKEN ENV variable is not defined")
	}

	return tokenCache
}
