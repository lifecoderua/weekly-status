package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"weekly-status/slack"
)

// os.Getenv("BOT_TOKEN")

// Route slack event
func Route(w http.ResponseWriter, r *http.Request) {
	var event SlackEvent
	switch r.URL.Path[0:] {
	case "/event/":
		log.Print("Event occured")
	case "/slack/daily":
		json.NewDecoder(r.Body).Decode(&event)
		// log.Printf("Daily payload: %s", event)
		routeEvent(w, r, event)
	case "/slack/weekly":
		log.Print("Weekly")
	default:
		echo(w, r)
	}
}

func routeEvent(w http.ResponseWriter, r *http.Request, event SlackEvent) {
	log.Printf("Daily payload: %s", event)
	switch event.Type {
	case "url_verification":
		fmt.Fprintf(w, "%s", event.Challenge)
	case "event_callback":
		log.Printf("Event with text received:\n %s\n", event.Event.Text)
		now := time.Now()
		isMonday := now.Weekday() == time.Monday

		// TODO: remove in prod
		DEBUG := true
		if DEBUG || isMonday {
			// isDailyReport :=
			// isWeeklyReport :=
			// default: ignore
			log.Printf("I am Monday, ask user or proxy his reply!")

			switch event.GetCommand() {
			case Init:
				// ToDo: channel data => daily_status => user name<=>channel
				// TODO: and make this one level up ^. isMonday only matters for DailyReceived
				log.Print("!!## Init")
				info := slack.GetChannelMembers("G71FJA4Q7")
				log.Printf("%s", info)
			case DailyReceived:
			case WeeklyReceived:
			}
			if event.Event.Text == "Log me" {
				recipient := "U69C2NL7N"
				message := "Thank you for debugging me"
				go slack.Message(recipient, message)
			}
		}
	default:
		log.Printf("!! Unexpected event type %s", event.Type)
	}
}

func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

	log.Printf("Caught input %s", r.URL.Path[1:])
}
