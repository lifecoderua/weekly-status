package app

// Route incoming event
// func Route() {
// 	log.Println("yay")
// }

// SlackEvent ...
type SlackEvent struct {
	Type      string
	Challenge string
	Token     string
	User      string
	Event     event
}

type event struct {
	Text string
}
