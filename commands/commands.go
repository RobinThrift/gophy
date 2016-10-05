package commands

import (
	"regexp"
	"github.com/nlopes/slack"
)

type Command struct {
	Name string
	pattern  *regexp.Regexp
	handler func(text string, user *slack.User) (error, string)
}

func (self *Command) IsApplicable(text string) bool {
	return self.pattern.MatchString(text)
}

func (self *Command) HandleMsg(text string, user *slack.User) (error, string) {
	return self.handler(text, user)
}
