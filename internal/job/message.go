package job

import (
	"math/rand"
	"strconv"
	"strings"
)

// TODO: Remplace this by CHATGPT one day
var templates []string = []string{
	`
Hey! 👋

Check out this article by _$AUTHOR$_ on _$PLATFORM$_ about *$TITLE$*.

If you're interested in $TAGS$ this is a must-read!

🔗 $LINK$
`,
	`
_$PLATFORM$_ just published an intriguing piece by _$AUTHOR$_.

The article highlights *$TITLE$*.

If you're into $TAGS$, give it a read!

🔗 $LINK$
`,
	`
I came across this fascinating 🤔 article by _$AUTHOR$_ on _$PLATFORM$_ about *$TITLE$*.

If you're curious about $TAGS$, take a look!

🔗 $LINK$
`,
	`
❗ Attention❗

Check out this article by _$AUTHOR$_ on _$PLATFORM$_ about *$TITLE$*.

Don't miss out! $TAGS$

🔗 $LINK$
`,
	`
Want to learn more 👨‍🏫 about *$TITLE$*?

_$AUTHOR$'s_ _$PLATFORM$_ article dives deep into $TAGS$.

Check it out!

🔗 $LINK$
`,
	`
Hey everyone 👋,

I found this interesting _$PLATFORM$_ article by _$AUTHOR$_ about *$TITLE$*.

$TAGS$

🔗 $LINK$
`, `
👀 If you're interested in $TAGS$,

you won't want to miss this _$PLATFORM$_ article by _$AUTHOR$_ about *$TITLE$*.

Check it out!

🔗 $LINK$
`, `
👀 💰 💡💡💡

👉🏼 📝👩🏻‍💼 _$AUTHOR$_ on _$PLATFORM$_ about *$TITLE$* for $TAGS$

🔗 $LINK$
`, `
*$TITLE$*,

_$AUTHOR$'s_ _$PLATFORM$_ article explores this topic $TAGS$.

If you're curious, give it a read!

🔗 $LINK$
`, `
_$PLATFORM$_ just published an insightful piece by _$AUTHOR$_ about *$TITLE$*.

Don't miss out on this one!

$TAGS$

🔗 $LINK$
`,
}

// Will inject the parameters into the template choose randomly.
func buildMessage(title string, platform string, author string, tags []string, link string) string {
	if len(tags) == 0 {
		tags = []string{"UNKNOWN"}
	}

	text := templates[rand.Intn(len(templates))]
	replacer := strings.NewReplacer("$TITLE$", title, "$AUTHOR$", author, "$PLATFORM$", platform, "$TAGS$", "#"+strings.Join(tags, ", #"), "$LINK$", link)

	return replacer.Replace(text)
}

var topicTemplates []string = []string{
	`
🚀 Exciting Discovery!

Dive into the latest buzz surrounding _$MAIN_SUBJECT$_! We've unearthed a treasure trove of articles – a total of $NUMBER_ARTICLES$ captivating reads. Explore the depths of knowledge with these compelling pieces:

$LINKS$
🔖 Quick Tags: $TAGS$

Ready to embark on this journey? Click the links and discover the fascinating world of _$MAIN_SUBJECT$_! 🌐✨
	`,
	`
🌟 Unlock the Secrets of _$MAIN_SUBJECT$_!

Embark on a thrilling exploration of _$MAIN_SUBJECT$_! We've curated $NUMBER_ARTICLES$ articles that promise to reveal the mysteries and wonders hidden within:

$LINKS$
🔖 Tags for Quick Insight: $TAGS$

Ready to be spellbound? Click the links and uncover the magic of _$MAIN_SUBJECT$_! 🌟✨
	`,
	`
🔍 Stay Ahead of the Curve with _$MAIN_SUBJECT$_ Breakthroughs!

Exciting times await! We've got $NUMBER_ARTICLES$ groundbreaking articles on the latest breakthroughs in _$MAIN_SUBJECT$_ . Don't miss out on the future – check out these insightful reads:

$LINKS$

🔖 Tags for a Quick Overview: $TAGS$

Ready to stay informed? Click the links and be part of the MAIN_SUBJECT revolution! 🚀🌐
	`,
}

func buildMessageFromTopic(topic Topic) string {

	nbArticles := len(topic.Articles)
	links := ""
	for k, article := range topic.Articles {
		links += "\n [Article " + strconv.Itoa(k+1) + "](" + article.Link + ")"
	}

	text := topicTemplates[rand.Intn(len(topicTemplates))]
	replacer := strings.NewReplacer(
		"$MAIN_SUBJECT$", topic.Subject,
		"$TAGS$", "#"+strings.Join(topic.Set.ToSlice(), ", #"),
		"$LINKS$", links,
		"$NUMBER_ARTICLES$", strconv.Itoa(nbArticles),
	)

	return replacer.Replace(text)
}
