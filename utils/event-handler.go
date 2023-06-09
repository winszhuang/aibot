package utils

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type EventHandler struct {
	Event  *linebot.Event
	Bot    *linebot.Client
	UserId string
}

func (h *EventHandler) SendText(text string) error {
	_, err := h.Bot.ReplyMessage(
		h.Event.ReplyToken,
		linebot.NewTextMessage(text),
	).Do()
	return err
}

func (h *EventHandler) SendImage(url string) error {
	_, err := h.Bot.ReplyMessage(
		h.Event.ReplyToken,
		linebot.NewImageMessage(url, url),
	).Do()
	return err
}
