package job

import (
	"fmt"
	"time"

	"github.com/Vico1993/Otto-client/otto"
	"github.com/Vico1993/otto-cron-chats/internal/service"
	"github.com/Vico1993/otto-cron-chats/internal/utils"
	"github.com/go-co-op/gocron"
)

var (
	chatsTag  = "chats"
	Scheduler = gocron.NewScheduler(time.UTC)
	telegram  = service.NewTelegramService()
)

func Main(client otto.Client) {
	chats := client.Chat.ListAll()

	if len(chats) == 0 {
		fmt.Println("No chats retrieved from API")
		return
	}

	// retreve all chats job in the scheduler
	jobs, err := Scheduler.FindJobsByTag(chatsTag)
	if err == nil {
		fmt.Println("Error retrieving chats job")
		return
	}

	// No job found but we have chats
	// OR if we have more or less chats than before
	// we need to start the worker
	if len(chats) != len(jobs) {
		fmt.Println("Starting the process with new list of chats")

		// First clean our current list of jobs
		err := Scheduler.RemoveByTag(chatsTag)
		if err != nil {
			fmt.Println("Couldn't reset chats")
		}

		// Creating the new jobs
		n := 1
		delay := utils.GetDelay(len(chats))
		for _, chat := range chats {
			// Copy val to be sure it's not overrited with the next iteration
			chat := chat

			// Start at different time to avoid parsing all feed at the same time
			when := delay * n
			fmt.Println("Adding Job -> " + chat.TelegramChatId + " threadId -> " + chat.TelegramThreadId)
			_, err := Scheduler.Every(2).
				Hour().
				Tag(chatsTag).
				StartAt(time.Now().Add(time.Duration(when) * time.Minute)).
				Do(func() {
					fmt.Println("Start : " + chat.TelegramChatId)

					articles := client.Chat.ListLatestArticles(chat.TelegramChatId, chat.TelegramThreadId)

					if len(articles) > 0 {
						notify(articles, chat)
					}

					// Update the time parsed
					client.Chat.UpdateParsedTime(chat.TelegramChatId, chat.TelegramThreadId)
					fmt.Println("End")
				})

			if err != nil {
				fmt.Println("Error initiate the cron for: " + chat.TelegramChatId + " - " + err.Error())
			}

			n += 1
		}
	}

	Scheduler.StartBlocking()
}
