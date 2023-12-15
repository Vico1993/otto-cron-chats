package job

import (
	"testing"

	"github.com/Vico1993/Otto-client/otto"
	"github.com/Vico1993/otto-cron-chats/internal/service"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
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

func TestFindTagFromTitle(t *testing.T) {
	sentence := "Microsoft CTO Kevin Scott on how AI and art will coexist in the future"

	result := findTagFromTitle(sentence)

	assert.True(t, result.Contains("kevin", "scott", "cto", "microsoft", "art", "future"))
}

func TestJaccardSimilarityDifferentSet(t *testing.T) {
	firstSet := mapset.NewSet[string]()
	firstSet.Append("microsoft", "game", "xbox", "blizzard")

	secondSet := mapset.NewSet[string]()
	secondSet.Append("bitcoin", "microsoft", "blockchain")

	assert.Lessf(t, jaccardSimilarity(firstSet, secondSet), 1.0, "error message %s", "formatted")
	assert.Lessf(t, jaccardSimilarity(firstSet, secondSet), 0.5, "error message %s", "formatted")
	assert.Greaterf(t, jaccardSimilarity(firstSet, secondSet), 0.0, "error message %s", "formatted")
}

func TestJaccardSimilaritySameTopic(t *testing.T) {
	firstSet := mapset.NewSet[string]()
	firstSet.Append("openai", "sam", "altman", "fired", "ceo")

	secondSet := mapset.NewSet[string]()
	secondSet.Append("sam", "openai", "altman", "removed", "ceo")

	assert.Lessf(t, jaccardSimilarity(firstSet, secondSet), 1.0, "error message %s", "formatted")
	assert.Greaterf(t, jaccardSimilarity(firstSet, secondSet), 0.5, "error message %s", "formatted")
	assert.Greaterf(t, jaccardSimilarity(firstSet, secondSet), 0.0, "error message %s", "formatted")
}

func TestMatch2ArticlesSameTopic(t *testing.T) {
	article1 := otto.Article{
		Title: "Sam Altman ousted as OpenAI’s CEO",
	}

	article2 := otto.Article{
		Title: "Sam Altman, OpenAI’s CEO, Is Ousted by Company’s Board",
	}

	res := match([]otto.Article{article1, article2})

	assert.Len(t, res, 1, "Should have only 1 topic")
	assert.Equal(t, res[0].Subject, article1.Title, "Subject should be equal to the title of the first article")
}

func TestMatch2ArticlesDifferentTopic(t *testing.T) {
	article1 := otto.Article{
		Title: "Sam Altman ousted as OpenAI’s CEO",
	}

	article2 := otto.Article{
		Title: "Microsoft CEO testifies that Google’s power in search is ubiquitous",
	}

	res := match([]otto.Article{article1, article2})

	assert.Len(t, res, 2, "Should have only 2 topics")
	assert.Equal(t, res[0].Subject, article1.Title, "Subject of topic 1 should be equal to the title of the first article")
	assert.Equal(t, res[1].Subject, article2.Title, "Subject of topic 2 should be equal to the title of the second article")
}
