package bot

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/wei840222/certchecker/db"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var Bot *tgbotapi.BotAPI

func init() {
	bot, err := tgbotapi.NewBotAPI("628484017:AAHw9EcMePoxGLBR-4fd8BnXYHKRVdSwUuo")
	if err != nil {
		panic(err)
	}
	bot.Debug = true
	Bot = bot
}

func HandleUpdate() {
	uc := tgbotapi.NewUpdate(0)
	uc.Timeout = 60
	updates, err := Bot.GetUpdatesChan(uc)
	if err != nil {
		panic(err)
	}
	for u := range updates {
		if u.Message == nil {
			continue
		}
		if u.Message.Text != "" {
			if err := handleText(u.Message.Chat.ID, u.Message.MessageID, u.Message.Text); err != nil {
				log.Printf("test message error: %s", err)
				msg := tgbotapi.NewMessage(u.Message.Chat.ID, err.Error())
				msg.ReplyToMessageID = u.Message.MessageID
				Bot.Send(msg)
			}
		}
	}
}

func handleText(chatID int64, messageID int, text string) error {
	if strings.HasPrefix(text, "/add") {
		if len(strings.Split(strings.TrimSpace(strings.TrimPrefix(text, "/add")), " ")) != 2 {
			return errors.New("invalid args")
		}
		name := strings.Split(strings.TrimSpace(strings.TrimPrefix(text, "/add")), " ")[0]
		domain := strings.Split(strings.TrimSpace(strings.TrimPrefix(text, "/add")), " ")[1]
		err := db.CreateDomain(&db.Domain{Name: name, Host: domain})
		if err != nil {
			return err
		}
		msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Create domain %s success!", domain))
		msg.ReplyToMessageID = messageID
		if _, err := Bot.Send(msg); err != nil {
			return err
		}
	}

	if strings.HasPrefix(text, "/list") {
		domains, err := db.ListDomain()
		if err != nil {
			return err
		}
		message := "Domains:\n"
		for _, d := range domains {
			message += fmt.Sprintf("ID: %d, Name: %s, Host: %s\n", d.ID, d.Name, d.Host)
		}
		msg := tgbotapi.NewMessage(chatID, message)
		msg.ReplyToMessageID = messageID
		if _, err := Bot.Send(msg); err != nil {
			return err
		}
	}

	if strings.HasPrefix(text, "/del") {
		id, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(text, "/del")))
		if err != nil {
			return err
		}
		if err := db.DeleteDomain(uint(id)); err != nil {
			return err
		}
		msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Delete domain %d success!", id))
		msg.ReplyToMessageID = messageID
		if _, err := Bot.Send(msg); err != nil {
			return err
		}
	}

	return nil
}