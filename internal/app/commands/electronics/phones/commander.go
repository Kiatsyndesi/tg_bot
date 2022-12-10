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
	default:
		c.Default(msg)
	}
}

/*
func (c *ElectronicsPhonesCommander) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("Value from recovered panic: %v", panicValue)
		}
	}()

	if update.CallbackQuery != nil {
		parsedData := CommandData{}

		json.Unmarshal([]byte(update.CallbackQuery.Data), &parsedData)

		msg := tgbotapi.NewMessage(
			update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Command: %+v\n", parsedData),
		)

		c.bot.Send(msg)

		return
	}

	if update.Message == nil {
		return
	}

	switch update.Message.Command() {
	case "help":
		c.Help(update.Message)
	case "list":
		c.List(update.Message)
	case "get":
		c.Get(update.Message)
	default:
		c.Default(update.Message)
	}
}
*/
