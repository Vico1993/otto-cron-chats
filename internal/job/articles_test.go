package job

import (
	"testing"

	"github.com/Vico1993/Otto-client/otto"
	"github.com/Vico1993/otto-cron-chats/internal/service"
	"github.com/google/uuid"
)

func TestFetch1Article(t *testing.T) {
	article := otto.Article{
		Id:     uuid.New().String(),
		FeedId: uuid.New().String(),
		Title:  "Super Title",
		Source: "Title",
		Author: "Unknown",
		Link:   "https://super.com/title",
		Tags:   []string{"tag1", "tag2"},
	}

	chat := otto.Chat{
		Id:               "12",
		TelegramChatId:   "12314",
		TelegramThreadId: "",
		Tags:             []string{"tag2"},
	}

	templates = []string{"MESSAGE"}

	mockTelegramService := new(service.MocksTelegramService)
	telegram = mockTelegramService
	mockTelegramService.On("TelegramUpdateTyping", chat.TelegramChatId, true).Return()
	mockTelegramService.On("TelegramUpdateTyping", chat.TelegramChatId, false).Return()
	mockTelegramService.On("TelegramPostMessage", chat.TelegramChatId, "", "MESSAGE").Return()

	fetch([]otto.Article{article}, chat)

	mockTelegramService.AssertCalled(t, "TelegramUpdateTyping", chat.TelegramChatId, true)
	mockTelegramService.AssertCalled(t, "TelegramUpdateTyping", chat.TelegramChatId, false)
	mockTelegramService.AssertCalled(t, "TelegramPostMessage", chat.TelegramChatId, "", "MESSAGE")
}

func TestFetch1ArticleButWithThreadId(t *testing.T) {
	article := otto.Article{
		Id:     uuid.New().String(),
		FeedId: uuid.New().String(),
		Title:  "Super Title",
		Source: "Title",
		Author: "Unknown",
		Link:   "https://super.com/title",
		Tags:   []string{"tag1", "tag2"},
	}

	threadId := "134"
	chat := otto.Chat{
		Id:               "12",
		TelegramChatId:   "12314",
		TelegramThreadId: threadId,
		Tags:             []string{"tag2"},
	}

	templates = []string{"MESSAGE"}

	mockTelegramService := new(service.MocksTelegramService)
	telegram = mockTelegramService
	mockTelegramService.On("TelegramUpdateTyping", chat.TelegramChatId, true).Return()
	mockTelegramService.On("TelegramUpdateTyping", chat.TelegramChatId, false).Return()
	mockTelegramService.On("TelegramPostMessage", chat.TelegramChatId, threadId, "MESSAGE").Return()

	fetch([]otto.Article{article}, chat)

	mockTelegramService.AssertCalled(t, "TelegramUpdateTyping", chat.TelegramChatId, true)
	mockTelegramService.AssertCalled(t, "TelegramUpdateTyping", chat.TelegramChatId, false)
	mockTelegramService.AssertCalled(t, "TelegramPostMessage", chat.TelegramChatId, threadId, "MESSAGE")
}
