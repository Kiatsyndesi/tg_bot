package phones

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *ElectronicsPhonesCommander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote this: "+inputMessage.Text)

	_, err := c.bot.Send(msg)

	if err != nil {
		log.Printf("ElectronicsPhonesCommander.Help: error sending reply message to chat - %v", err)
	}
}
