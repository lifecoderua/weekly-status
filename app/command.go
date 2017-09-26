package app

// Command received from Slack chat
type Command int

// _
const (
	None Command = iota
	Init
	DailyReceived
	WeeklyReceived
)
