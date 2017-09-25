package app

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
