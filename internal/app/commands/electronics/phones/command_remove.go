package phones

import (
	"github.com/Kiatsyndesi/tg_bot/internal/service/electronics/phones"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (c *ElectronicsPhonesCommander) Remove(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	id, err := strconv.Atoi(args)

	if err != nil {
		log.Printf("wrong args in remove command")

		return
	}

	phones.AllPhones, err = c.phonesService.Remove(uint64(id))

	if err != nil {
		log.Print(err)

		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Success",
	)

	_, err = c.bot.Send(msg)

	if err != nil {
		log.Printf("fail to remove product with id %d: %v", id, err)

		return
	}
}
