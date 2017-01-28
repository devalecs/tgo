package main

import "gopkg.in/devalecs/tgo.v1"

func main() {
	c := tgo.NewClient("yourTelegramBotAPIToken")

	updatesChan := c.GetUpdatesChan(tgo.GetUpdatesParams{
		Timeout: 60,
	})

	for update := range updatesChan {
		if update.Message == nil {
			continue
		}

		c.SendMessage(tgo.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Pong",
		})
	}
}
