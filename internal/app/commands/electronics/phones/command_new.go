package phones

import (
	"github.com/Kiatsyndesi/tg_bot/internal/service/electronics/phones"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *ElectronicsPhonesCommander) New(inputMessage *tgbotapi.Message) {
	arg := inputMessage.CommandArguments()

	newPhone := phones.Phone{
		ID:    c.phonesService.GetLastID() + 1,
		Title: arg,
	}

	var err error
	phones.AllPhones, err = c.phonesService.New(newPhone)

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
		log.Printf("fail to add product %+v: %v", newPhone, err)

		return
	}
}
