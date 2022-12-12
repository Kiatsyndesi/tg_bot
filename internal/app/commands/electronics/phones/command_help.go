package phones

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *ElectronicsPhonesCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help_electronics_phones - this is my help for you\n"+
			"/list_electronics_phones - list of products\n"+
			"/get_electronics_phones 1 - find phone by index\n"+
			"/remove_electronics_phones 1 - remove phone from list by index",
	)

	_, err := c.bot.Send(msg)

	if err != nil {
		log.Printf("ElectronicsPhonesCommander.Help: error sending reply message to chat - %v", err)
	}
}
