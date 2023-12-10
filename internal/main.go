package main

import (
	"os"

	"github.com/Vico1993/Otto-client/otto"
	"github.com/Vico1993/otto-cron-chats/internal/job"
	"github.com/Vico1993/otto-cron-chats/internal/service"
	"github.com/Vico1993/otto-cron-chats/internal/utils"
	"github.com/subosito/gotenv"
)

func main() {
	// load .env file if any otherwise use env set
	_ = gotenv.Load()

	OttoClient := otto.NewClient(
		nil,
		os.Getenv("OTTO_API_URL"),
	)

	// Notify update if chat present
	if os.Getenv("TELEGRAM_ADMIN_CHAT_ID") != "" {
		version := utils.RetrieveVersion()

		service.NewTelegramService().TelegramPostMessage(
			os.Getenv("TELEGRAM_ADMIN_CHAT_ID"),
			"",
			`🚀 🚀 [CRON-CHATS] Version: *`+version+`* Succesfully deployed . 🚀 🚀`,
		)
	}

	job.Main(*OttoClient)
}
