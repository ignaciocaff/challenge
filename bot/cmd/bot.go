package main

import (
	"botjobsity/env"
	"botjobsity/services"
)

func main() {
	env := env.GetEnv(".env")
	bot := services.NewBot(env)
	bot.Connect()
	<-bot.Forever
}
