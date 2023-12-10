package job

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/Vico1993/Otto-client/otto"
)

// Fetch all articles
func fetch(articles []otto.Article, chat otto.Chat) {
	fmt.Println("Articles found:" + strconv.Itoa(len(articles)))

	telegram.TelegramUpdateTyping(chat.TelegramChatId, true)
	for _, article := range articles {
		article := article

		host := article.Source
		u, err := url.Parse(article.Source)
		if err == nil {
			host = u.Host
		}

		fmt.Println("send", host)
		telegram.TelegramPostMessage(
			chat.TelegramChatId,
			chat.TelegramThreadId,
			buildMessage(
				article.Title,
				host,
				article.Author,
				article.Tags,
				article.Link,
			),
		)
	}
	telegram.TelegramUpdateTyping(chat.TelegramChatId, false)
}
