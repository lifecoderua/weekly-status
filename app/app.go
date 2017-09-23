package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	default:
		log.Printf("!! Unexpected event type %s", event.Type)
	}
}

func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

	log.Printf("Caught input %s", r.URL.Path[1:])
}
