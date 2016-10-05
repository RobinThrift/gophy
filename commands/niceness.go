package commands


import (
	"fmt"
	"regexp"
	"github.com/nlopes/slack"
)

var nicePersonPattern *regexp.Regexp = regexp.MustCompile(`\s*[iI]s (@?[a-zA-Z0-9]+)(?: a)? nice(?: person)?\??`)

func isPersonNiceHandler(text string, user *slack.User) (error, string) {
	userName := nicePersonPattern.FindStringSubmatch(text)[1]
	modifier := "*NOT* "
	if userName == "robin" {
		modifier = ""
	}
	return nil, fmt.Sprintf("I can say, without a shadow of a doubt, that %s is, in fact, %sa nice person!", userName, modifier)
}

var IsPersonNiceCmd = Command{
	Name: "IsPersonNice",
	pattern: nicePersonPattern,
	handler: isPersonNiceHandler,
}
