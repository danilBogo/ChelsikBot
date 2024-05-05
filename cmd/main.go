package main

import (
	"ChelsikBot/internal/app"
	"ChelsikBot/internal/app/health"
)

func main() {
	application := app.NewApp()
	application.Start()
	health.Start(application)
}
