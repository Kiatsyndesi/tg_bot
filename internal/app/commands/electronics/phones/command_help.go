package phones

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *ElectronicsPhonesCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - this is my help for you\n"+
			"/list - list of products\n"+
		"/get - find phone by argument",
	)

	_, err := c.bot.Send(msg)

	if err != nil {
		log.Printf("ElectronicsPhonesCommander.Help: error sending reply message to chat - %v", err)
	}
}
