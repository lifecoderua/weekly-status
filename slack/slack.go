package slack

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

const api = "https://slack.com/api/"

var tokenCache string

type slackPayload struct {
	token   string
	channel string
	text    string
	asUser  bool
}

// TODO: Getter or a clean way to use same space for channels and groups
// ChannelInfo know information about your channel
type ChannelInfo struct {
	Group ChannelInfoGroup
}

// ChannelInfoGroup know information about your channel Group
type ChannelInfoGroup struct {
	id      string
	name    string
	members []string
}

// Message recipient
func Message(channel string, message string) {
	log.Print("MSG<><><>")
	// run("chat.postMessage", channel, message)
	run("chat.postMessage", slackPayload{channel: channel, text: message})
}

// GetChannelMembers return channel members information
func GetChannelMembers(channel string) ChannelInfo {
	var info ChannelInfo
	endpoint := "channels.info"
	if channel[:1] == "G" {
		endpoint = "groups.info"
	}

	res := run(endpoint, slackPayload{channel: channel})
	json.NewDecoder(res.Body).Decode(&info)
	return info
}

// def room_members(channel)
// endpoint = 'G' === channel[0] ? 'groups.info' : 'channels.info'

// run(endpoint, {
// 	channel: channel
// })['group']['members']
// end

// func run(endpoint string, channel string, text string) {
func run(endpoint string, payload slackPayload) *http.Response {
	uri := api + endpoint

	resp, err := http.PostForm(uri,
		url.Values{
			"token":   {token()},
			"channel": {payload.channel},
			"text":    {payload.text},
			"as_user": {"true"}})

	if nil != err {
		fmt.Println("errorination happened reading the body", err)
		return &http.Response{}
	}

	defer resp.Body.Close()

	return resp
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
