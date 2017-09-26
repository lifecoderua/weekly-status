package slack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	user    string
	asUser  bool
}

// ChannelInfo know information about your channel
// TODO: Getter or a clean way to use same space for channels and groups
type ChannelInfo struct {
	Group ChannelInfoGroup
}

// ChannelInfoGroup know information about your channel Group
type ChannelInfoGroup struct {
	ID      string
	Name    string
	Members []string
}

// MemberWrap is a chat user wrapper
type MemberWrap struct {
	User Member
}

// Member is a chat user
type Member struct {
	ID    string
	Name  string
	IsBot bool `json:"is_bot"`
}

// Message recipient
func Message(channel string, message string) {
	log.Print("MSG<><><>")
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

	errx := json.NewDecoder(res.Body).Decode(&info)
	log.Printf("%s )) %s", errx, info)

	mem0 := getMemberInfo(info.Group.Members[0])
	log.Printf("%s << ", mem0)
	return info
}

func getMemberInfo(user string) Member {
	var info MemberWrap

	res := run("users.info", slackPayload{user: user})
	body, _ := ioutil.ReadAll(res.Body)
	log.Printf("%s, %s", user, body)
	errx := json.NewDecoder(res.Body).Decode(&info)
	log.Printf("%s )) %s", errx, info.User)
	return info.User
}

// def room_members(channel)
// endpoint = 'G' === channel[0] ? 'groups.info' : 'channels.info'

// run(endpoint, {
// 	channel: channel
// })['group']['members']
// end

// func run(endpoint string, channel string, text string) {
// func run(endpoint string, payload slackPayload) *http.Response {
func run(endpoint string, payload map[string][]string) *http.Response {
	uri := api + endpoint

	payload["token"] = token()
	payload["as_user"] = "true"

	resp, err := http.PostForm(uri,
		url.Values(payload))

	// resp, err := http.PostForm(uri,
	// 	url.Values{
	// 		"token":   {token()},
	// 		"channel": {payload.channel},
	// 		"text":    {payload.text},
	// 		"as_user": {"true"}})

	if nil != err {
		fmt.Println("errorination happened reading the body", err)
		return &http.Response{}
	}

	// defer resp.Body.Close()

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
