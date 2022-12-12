package phones

import (
	"github.com/Kiatsyndesi/tg_bot/internal/app/path"
	"github.com/Kiatsyndesi/tg_bot/internal/service/electronics/phones"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type ElectronicsPhonesCommander struct {
	bot           *tgbotapi.BotAPI
	phonesService *phones.Service
}

type CommandData struct {
	Offset string `json:"offset"`
}

func NewElectronicsPhonesCommander(bot *tgbotapi.BotAPI) *ElectronicsPhonesCommander {
	phonesService := phones.NewService()

	return &ElectronicsPhonesCommander{
		bot:           bot,
		phonesService: phonesService,
	}
}

func (c *ElectronicsPhonesCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("ElectronicsPhonesCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *ElectronicsPhonesCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "remove":
		c.Remove(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	default:
		//bot uses router which has own non-command msg handling. Code below for use phoneCommander directly
		c.Default(msg)
	}
}
