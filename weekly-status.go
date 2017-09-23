package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	//"encoding/json"

	"github.com/joho/godotenv"
)

// import _ "github.com/joho/godotenv"

// import _ "github.com/joho/godotenv/autoload"

// SlackEvent ...
type SlackEvent struct {
	Type  string
	Token string
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

	log.Printf("Caught input %s", r.URL.Path[1:])
}

func route(w http.ResponseWriter, r *http.Request) {
	var event SlackEvent
	switch r.URL.Path[0:] {
	case "/event/":
		log.Print("Event occured")
	case "/slack/daily":
		json.NewDecoder(r.Body).Decode(&event)
		log.Printf("Daily payload: %s", event)
	case "/slack/weekly":
		log.Print("Weekly")
	default:
		handler(w, r)
	}
}

func main() {
	_ = godotenv.Load()

	log.Printf("Bot token: %s", os.Getenv("BOT_TOKEN"))

	http.HandleFunc("/", route)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
