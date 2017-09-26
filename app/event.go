package app

import (
	"regexp"
)

// SlackEvent ...
type SlackEvent struct {
	Type      string
	Challenge string
	Token     string
	Event     event
}

type event struct {
	Text    string
	User    string
	Channel string
}

// GetCommand evaluates which command was issued from the chat
func (event *SlackEvent) GetCommand() Command {
	rInit := regexp.MustCompile("^weekly!init")
	if rInit.MatchString(event.Event.Text) {
		return Init
	}

	return None
}
