package main

import (
	"github.com/MlDenis/go-url-shortener/internal/app"
	"log"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatal("Server startup error", err)
	}
}
