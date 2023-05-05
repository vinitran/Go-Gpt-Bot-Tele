package bot

import (
	"github.com/urfave/cli/v2"
	"os"
)

func NewCommand() *cli.Command {
	return &cli.Command{
		Name:  "bot",
		Usage: "Start the API server",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			return startBot()
		},
		Before: func(c *cli.Context) error {
			return beforeBot()
		},
	}
}

func beforeBot() error {
	return nil
}

func startBot() error {
	bot, err := NewBot(os.Getenv(envBotToken))
	if err != nil {
		return err
	}

	err = bot.Handler()
	return err
}
