package main

import (
	ai "aibot/services"
	. "aibot/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/line/line-bot-sdk-go/v7/linebot/httphandler"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	handler, err := httphandler.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}
	bot, err := handler.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/callback", handler)
	handler.HandleEvents(func(events []*linebot.Event, r *http.Request) {
		for _, event := range events {
			userId := event.Source.UserID
			fmt.Println(userId)

			eventHandler := &EventHandler{Event: event, Bot: bot, UserId: userId}

			switch event.Type {
			case linebot.EventTypeMessage:
				handleMessage(eventHandler)
			}
		}
	})

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}

func handleMessage(eh *EventHandler) {
	switch messageData := eh.Event.Message.(type) {
	case *linebot.TextMessage:
		message := strings.TrimSpace(messageData.Text)
		fmt.Println(message)

		err := eh.SendText(ai.Reply(message))
		if err != nil {
			log.Fatal(err)
		}
	}
}
