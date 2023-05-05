package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"gptBot/cmd/bot"
	"log"
	"os"
)

const (
	envPath = ".env"
)

func init() {
	if err := godotenv.Overload(envPath); err != nil {
		log.Fatal("Load env error", err.Error())
		return
	}
}

func main() {
	app := &cli.App{
		Name:  "GPT-Bot",
		Usage: "GPT Bot",
		Action: func(*cli.Context) error {
			fmt.Println("use --help")
			return nil
		},
		Flags: []cli.Flag{},
		Commands: []*cli.Command{
			bot.NewCommand(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("[Main] Run CLI error:", err.Error())
		return
	}
}
