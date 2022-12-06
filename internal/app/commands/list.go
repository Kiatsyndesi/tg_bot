package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here are your products: \n\n"

	productList := c.productService.List()

	for _, product := range productList {
		outputMsgText += product.Title
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	c.bot.Send(msg)
}
