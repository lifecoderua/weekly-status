package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"weekly-status/app"
)

func main() {
	_ = godotenv.Load()

	http.HandleFunc("/", app.Route)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
