package job

import (
	"testing"

	"github.com/Vico1993/Otto-client/otto"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"
)

func TestBuildMessage(t *testing.T) {
	// Set up test data
	title := "Title"
	platform := "platform"
	author := "author"
	tags := []string{"tag1", "tag2"}
	link := "https://example.com"

	// Override templates
	templates = []string{"$TITLE$ on $PLATFORM$. by $AUTHOR$. check link: $LINK$ is about: $TAGS$"}

	res := buildMessage(title, platform, author, tags, link)
	assert.Equal(
		t,
		"Title on platform. by author. check link: https://example.com is about: #tag1, #tag2",
		res,
		"The template is not matching the expected translation!",
	)
}

func TestNoTags(t *testing.T) {
	// Set up test data
	title := "Title"
	platform := "platform"
	author := "author"
	tags := []string{}
	link := "https://example.com"

	// Override templates
	templates = []string{"$TITLE$ on $PLATFORM$. by $AUTHOR$. check link: $LINK$ is about: $TAGS$"}

	res := buildMessage(title, platform, author, tags, link)
	assert.Equal(
		t,
		"Title on platform. by author. check link: https://example.com is about: #UNKNOWN",
		res,
		"The template is not matching the expected translation!",
	)
}

func TestBuildMessageFromTopic(t *testing.T) {
	set := mapset.NewSet[string]()
	set.Add("foo")
	topic := Topic{
		Subject: "Super Foo",
		Set:     set,
		Articles: []otto.Article{
			{
				Link: "https://google.com",
			},
		},
	}

	// Override templates
	topicTemplates = []string{"$MAIN_SUBJECT$ on links: $LINKS$ is about: $TAGS$, with nbArticles: $NUMBER_ARTICLES$"}

	res := buildMessageFromTopic(topic)

	assert.Equal(
		t,
		"Super Foo on links: \n [Article 1](https://google.com) is about: #foo, with nbArticles: 1",
		res,
		"The template is not matching the expected translation!",
	)
}
