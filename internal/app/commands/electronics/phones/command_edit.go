package phones

import (
	"github.com/Kiatsyndesi/tg_bot/internal/service/electronics/phones"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

func (c *ElectronicsPhonesCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	phoneArgs := strings.SplitN(args, " TO ", 2)

	if len(phoneArgs) != 2 {
		log.Printf("Wrong phones for args: %+v", phoneArgs)

		return
	}

	phoneToEdit := phones.Phone{Title: phoneArgs[0]}
	updPhone := phones.Phone{Title: phoneArgs[1]}

	var err error
	phones.AllPhones, err = c.phonesService.Edit(phoneToEdit, updPhone)

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
		log.Printf("fail to edit product %+v to this %+v: error %v", phoneToEdit, updPhone, err)

		return
	}
}
