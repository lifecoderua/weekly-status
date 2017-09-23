package app

// Route incoming event
// func Route() {
// 	log.Println("yay")
// }

// SlackEvent ...
type SlackEvent struct {
	Type  string
	Token string
	User  string
	Text  string
}
