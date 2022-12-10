package phones

import (
	"encoding/json"
	"fmt"
	"github.com/Kiatsyndesi/tg_bot/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type CallbackListData struct {
	Offset string `json:"offset"`
}

func (c *ElectronicsPhonesCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}

	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)

	if err != nil {
		log.Printf("ElectronicsPhonesCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, fmt.Sprintf("Parsed: %v", parsedData))

	_, err = c.bot.Send(msg)

	if err != nil {
		log.Printf("ElectronicsPhoneCommander.CallbackList: error sending reply message to chat - %v", err)
	}
}
