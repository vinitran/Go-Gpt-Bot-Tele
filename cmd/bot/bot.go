package bot

import (
	tb "gopkg.in/telebot.v3"
	"gptBot/app/gpt"
	"os"
	"time"
)

type Bot struct {
	B      *tb.Bot `json:"b"`
	ChatId int64   `json:"chat_id"`
}

func NewBot(token string) (*Bot, error) {
	pref := tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	}

	tmpBot, err := tb.NewBot(pref)
	if err != nil {
		return nil, err
	}

	chatId, err := Int64(os.Getenv(envChatId))
	if err != nil {
		return nil, err
	}

	bot := &Bot{
		B:      tmpBot,
		ChatId: chatId,
	}

	return bot, nil
}

func (b *Bot) Handler() error {
	b.B.Handle(startCommand, func(c tb.Context) error {
		if b.AuthRequire(c) == false {
			return nil
		}

		err := c.Reply("Hello, I am an AI bot designed to answer your questions. Feel free to ask me anything you want and I will do my best to provide you with a helpful and informative response. Let's get started!")
		if err != nil {
			return err
		}

		return nil
	})

	b.B.Handle(questionCommand, func(c tb.Context) error {
		if b.AuthRequire(c) == false {
			return nil
		}

		if c.Message().Payload == "" {
			err := c.Reply("Tell me your question")
			if err != nil {
				return err
			}
			return nil
		}

		err := c.Reply("...")
		if err != nil {
			return err
		}
		answer, err := gpt.Predict(c.Message().Payload)
		if err != nil {
			return err
		}

		err = c.Reply(answer)
		if err != nil {
			return err
		}

		return nil
	})

	b.B.Start()
	return nil
}

func (b *Bot) AuthRequire(c tb.Context) bool {
	if c.Message().Chat.ID != b.ChatId {
		c.Send("You are not authorized to use this bot")
		return false
	}

	return true
}
