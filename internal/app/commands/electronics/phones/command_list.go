package phones

import (
	"encoding/json"
	"github.com/Kiatsyndesi/tg_bot/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *ElectronicsPhonesCommander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here are your phones: \n\n"

	phones := c.phonesService.List()

	for _, phone := range phones {
		outputMsgText += phone.Title
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	//hardcode data for callback
	serializedData, _ := json.Marshal(CallbackListData{
		Offset: "10",
	})

	callbackPath := path.CallbackPath{
		Domain:       "electronics",
		Subdomain:    "phones",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next Page", callbackPath.String()),
		),
	)

	_, err := c.bot.Send(msg)

	if err != nil {
		log.Printf("ElectronicsPhonesCommander.List: error sending reply message to chat - %v", err)
	}
}
