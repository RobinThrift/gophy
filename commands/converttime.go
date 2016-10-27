package commands

import (
	"github.com/nlopes/slack"
	"time"
	"regexp"
)

var convertTimePattern *regexp.Regexp = regexp.MustCompile("([\\d:]*) ([A-Za-z\\+\\d/_]+) in ([A-Za-z/_]+)")

func convertTimeHandler(text string, user *slack.User) (error, string) {
	parts := convertTimePattern.FindStringSubmatch(text)

	srcLoc, _ := time.LoadLocation(parts[2])

	const format = "15:04"
	t, err := time.ParseInLocation(format, parts[1], srcLoc)
	if err != nil {
		return err, ""
	}

	targetLoc, err := time.LoadLocation(parts[3])
	if err != nil {
		return err, ""
	}

	convT := t.In(targetLoc)

	return nil, convT.Format("15:04")
}

var ConvertTimeCommand = Command{
	Name: "ConvertTime",
	pattern: convertTimePattern,
	handler: convertTimeHandler,
}
