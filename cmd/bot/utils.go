package bot

import (
	"strconv"
)

const (
	startCommand    = "/start"
	questionCommand = "/question"
)

const (
	envBotToken = "BOT_TOKEN"
	envChatId   = "CHAT_ID"
)

func Int64(data string) (int64, error) {
	dataInt64, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		return 0, err
	}
	return dataInt64, nil
}
