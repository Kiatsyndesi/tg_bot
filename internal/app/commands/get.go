package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	arg, err := strconv.Atoi(args)

	if err != nil {
		log.Printf("Something wrong with argument: %v, error: %v", arg, err)

		return
	}

	product, err := c.productService.Get(arg)

	if err != nil {
		log.Printf("Something wrong to get from products: %v, error: %v", product, err)

		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		product.Title,
	)

	c.bot.Send(msg)
}
