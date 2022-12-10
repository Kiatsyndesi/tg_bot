package phones

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (c *ElectronicsPhonesCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)

	if err != nil {
		log.Println("wrong args", args)
		return
	}

	phone, err := c.phonesService.Get(idx)

	if err != nil {
		log.Printf("fail to get product with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		phone.Title,
	)

	_, err = c.bot.Send(msg)

	if err != nil {
		log.Printf("fail to get product with idx %d: %v", idx, err)
		return
	}
}
