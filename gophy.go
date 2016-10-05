package main

import (
	"fmt"
	"github.com/RobinThrift/gophy/commands"
	"github.com/nlopes/slack"
	"os"
)

func main() {
	cmds := []commands.Command{
		commands.RandomIntCommand,
	}

	api := slack.New(os.Getenv("SLACK_TOKEN"))

	rtm := api.NewRTM()

	go rtm.ManageConnection()

	var users []slack.User

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch evt := msg.Data.(type) {
			case *slack.HelloEvent:
				// @TODO: ignore for now, add logging later
			case *slack.ConnectedEvent:
				// @TODO: add logging
				users = evt.Info.Users
			case *slack.MessageEvent:
				// @TODO: add handling of messages here
				for _, cmd := range cmds {
					if cmd.IsApplicable(evt.Text) {
						err, msg := cmd.HandleMsg(evt.Text, findUserByID(users, evt.User))
						if err != nil {
							handleError(err.Error())
						} else {
							rtm.SendMessage(rtm.NewOutgoingMessage(msg, evt.Channel))
						}
					}
				}
			case *slack.RTMError:
				handleError(evt.Error())
				// @TODO: ignore for now, add logging later
			case *slack.InvalidAuthEvent:
				// @TODO: ignore for now, add logging later
				fmt.Printf("Invalid credentials")
				break Loop
			default:
			}
		}
	}
}

func findUserByID(users []slack.User, id string) *slack.User {
	for _, u := range users {
		if u.ID == id {
			return &u
		}
	}
	return nil
}

func handleError(err string) {
	fmt.Printf("Error: %s\n", err)
}
