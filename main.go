package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func goDotEnv(key string) string {
	err := godotenv.Load(".env")
	
	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv(key)
}

func main() {
	key := goDotEnv("SLACK_BOT_KEY")
	api := slack.New(key)

	channelID, timestamp, e := api.PostMessage(
		"C05U58K6EH0",
		slack.MsgOptionText("Hello world", false),
	)

	if e != nil {
		log.Fatal(e)
	}

	fmt.Printf("message sent successfully to channel at %s at %s", channelID, timestamp)
}