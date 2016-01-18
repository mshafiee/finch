package finchcommandstats

import (
	"bytes"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/syfaro/finch"
	"strconv"
)

type UserMessageCount map[string]int

var userMessages UserMessageCount

func init() {
	finch.RegisterCommand(&infoCollector{})
	finch.RegisterCommand(&infoCommand{})
	userMessages = make(UserMessageCount)
}

type infoCommand struct {
	finch.CommandBase
}

func (cmd *infoCommand) Help() finch.Help {
	return finch.Help{
		Name:        "Stats",
		Description: "Displays some statistics",
		Example:     "/stats@@",
		Botfather: [][]string{
			[]string{"stats", "Displays some statistics about bot usage"},
		},
	}
}

func (cmd *infoCommand) ShouldExecute(update tgbotapi.Update) bool {
	return finch.SimpleCommand("stats", update.Message.Text)
}

func (cmd *infoCommand) Execute(update tgbotapi.Update) error {
	b := &bytes.Buffer{}

	b.WriteString("Users seen\n\n")

	for userName, count := range userMessages {
		b.WriteString(userName)
		b.WriteString(" - ")
		b.WriteString(strconv.Itoa(count))
		b.WriteString("\n")
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, b.String())
	msg.ReplyToMessageID = update.Message.MessageID

	return cmd.Finch.SendMessage(msg)
}

type infoCollector struct {
	finch.CommandBase
}

func (cmd *infoCollector) Init(c *finch.CommandState, f *finch.Finch) error {
	cmd.CommandState = c
	cmd.Finch = f

	stored := cmd.Get("stats")
	if stored == nil {
		userMessages = make(UserMessageCount)
	} else {
		for user, count := range stored.(map[string]interface{}) {
			userMessages[user] = int(count.(float64))
		}
	}

	return nil
}

func (cmd *infoCollector) Help() finch.Help {
	return finch.Help{Name: "Stats Collector"}
}

func (cmd *infoCollector) ShouldExecute(update tgbotapi.Update) bool {
	return true
}

func (cmd *infoCollector) Execute(update tgbotapi.Update) error {
	if _, ok := userMessages[update.Message.From.String()]; !ok {
		userMessages[update.Message.From.String()] = 1
	} else {
		userMessages[update.Message.From.String()] += 1
	}

	cmd.CommandBase.Set("stats", userMessages)

	return nil
}
