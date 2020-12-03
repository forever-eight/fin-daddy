package main

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/forever-eight/fin-daddy/cmd/fin-daddy/currency"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("1436318012:AAF6mfxveYh213kd5Ge9ce_uydO3IqDjtuU")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		var msgText string

		if strings.ToLower(update.Message.Text) == "ты хороший папа" {
			msgText = "А ты хорошая дочь 🥰"
		} else {
			Name, Value := currency.GetCurrency(update.Message.Text)
			if Name == "" {
				msgText = "Я не знаю такой код валюты"
			} else {
				msgText = fmt.Sprintf("%s: %.2f₽", Name, Value)
			}
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
		_, err = bot.Send(msg)
		if err != nil {
			log.Print(err)
		}
	}
}
