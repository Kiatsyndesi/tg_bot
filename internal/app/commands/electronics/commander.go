package electronics

import (
	"github.com/Kiatsyndesi/tg_bot/internal/app/commands/electronics/phones"
	"github.com/Kiatsyndesi/tg_bot/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type ElectronicsCommander struct {
	bot *tgbotapi.BotAPI
	phonesCommander Commander
}

func NewElectronicsCommander(bot *tgbotapi.BotAPI) *ElectronicsCommander {
	return &ElectronicsCommander{
		bot: bot,
		phonesCommander: phones.NewElectronicsPhonesCommander(bot),
	}
}

func (c *ElectronicsCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Phones {
	case "Phones":
		c.phonesCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("ElectronicsCommander.HandleCommand: unknown subdomain - %s", callbackPath.Electronics)
	}
}

func (c *ElectronicsCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Phones {
	case "phones":
		c.phonesCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("ElectronicsCommander.HandleCommand: unknown subdomain - %s", commandPath.Phones)
	}
}
