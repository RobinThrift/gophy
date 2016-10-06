package commands

import (
	"fmt"
	"github.com/nlopes/slack"
	"math/rand"
	"regexp"
	"strconv"
)

var pattern *regexp.Regexp = regexp.MustCompile("randInt (-?[0-9]+) (-?[0-9]+)")

func randomIntHandler(text string, user *slack.User) (error, string) {
	intStrs := pattern.FindStringSubmatch(text)

	// @TODO: this could be nicer
	min, err := strconv.Atoi(intStrs[1])
	if err != nil {
		return err, ""
	}

	// @TODO: this could be nicer
	max, err := strconv.Atoi(intStrs[2])
	if err != nil {
		return err, ""
	}

	// @TODO: maybe don't hardcode messages?
	if min >= max {
		return nil, "Error, the second argument must be greater than the first!"
	}

	// @TODO: maybe don't hardcode messages?
	return nil, fmt.Sprintf("Your random integer, good sir/madam: %d", rand.Intn(max-min)+min)
}

var RandomIntCommand = Command{
	Name:    "RandomInt",
	pattern: pattern,
	handler: randomIntHandler,
}
