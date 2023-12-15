package job

import (
	"net/url"
	"strings"

	textrank "github.com/DavidBelicza/TextRank/v2"
	"github.com/Vico1993/Otto-client/otto"
	mapset "github.com/deckarep/golang-set/v2"
)

var (
	// Default Rule for parsing.
	rule = textrank.NewDefaultRule()
	// Default Language for filtering stop words.
	language = textrank.NewDefaultLanguage()
	// Default algorithm for ranking text.
	algorithmDef = textrank.NewDefaultAlgorithm()
	// Score needed to merge articles together
	similarityScore = 0.5
)

type Topic struct {
	Subject  string
	Set      mapset.Set[string]
	Articles []otto.Article
}

// Match articles together
func match(articles []otto.Article) []Topic {
	matchs := []Topic{}
	for _, article := range articles {
		set := findTagFromTitle(article.Title)

		// first iteration
		if len(matchs) > 0 {
			matched := false

			for key, match := range matchs {
				similarity := jaccardSimilarity(match.Set, set)

				if similarity >= similarityScore {
					matchs[key].Articles = append(matchs[key].Articles, article)
					matchs[key].Set = matchs[key].Set.Union(set)
					matched = true
				} else {
					continue
				}
			}

			if matched {
				continue
			}
		}

		matchs = append(matchs, Topic{
			Subject:  article.Title,
			Set:      set,
			Articles: []otto.Article{article},
		})
		continue
	}

	return matchs
}

// Notify user based on list of topic
func notify(articles []otto.Article, chat otto.Chat) {
	topics := match(articles)

	telegram.TelegramUpdateTyping(chat.TelegramChatId, true)
	for _, topic := range topics {
		topic := topic

		message := buildMessageFromTopic(topic)
		if len(topic.Articles) == 1 {
			article := topic.Articles[0]
			host := article.Source
			u, err := url.Parse(article.Source)
			if err == nil {
				host = u.Host
			}

			message = buildMessage(
				article.Title,
				host,
				article.Author,
				article.Tags,
				article.Link,
			)
		}

		telegram.TelegramPostMessage(
			chat.TelegramChatId,
			chat.TelegramThreadId,
			message,
		)
	}
	telegram.TelegramUpdateTyping(chat.TelegramChatId, false)
}

// JaccardSimilarity, as known as the Jaccard Index, compares the similarity of sample sets.
// This doesn't measure similarity between texts, but if regarding a text as bag-of-word,
// it can apply.
func jaccardSimilarity(s1set, s2set mapset.Set[string]) float64 {
	s1ands2 := s1set.Intersect(s2set).Cardinality()
	s1ors2 := s1set.Union(s2set).Cardinality()
	return float64(s1ands2) / float64(s1ors2)
}

// Extract important word from the title
func findTagFromTitle(title string) mapset.Set[string] {
	// TextRank object
	tr := textrank.NewTextRank()
	// Add text.
	tr.Populate(title, language, rule)
	// Run the ranking.
	tr.Ranking(algorithmDef)

	// Get all words order by weight.
	words := textrank.FindSingleWords(tr)

	set := mapset.NewSet[string]()
	for _, word := range words {
		set.Add(strings.ToLower(word.Word))
	}

	return set
}
