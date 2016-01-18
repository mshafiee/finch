package finchcommandhelp

import (
	"bytes"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/syfaro/finch"
)

func init() {
	finch.RegisterCommand(&helpCommand{})
}

type helpCommand struct {
	finch.CommandBase
}

func (cmd *helpCommand) Help() finch.Help {
	return finch.Help{
		Name:        "Help",
		Description: "Displays loaded commands and their help text",
		Example:     "/help@@",
		Botfather: [][]string{
			[]string{"help", "Displays available commands and help information"},
		},
	}
}

func (cmd *helpCommand) ShouldExecute(update tgbotapi.Update) bool {
	return finch.SimpleCommand("help", update.Message.Text)
}

func (cmd *helpCommand) Execute(update tgbotapi.Update) error {
	b := &bytes.Buffer{}

	if update.Message.CommandArguments() == "botfather" {
		for k, command := range cmd.Finch.Commands {
			help := command.Command.Help().BotfatherString()

			if help != "" {
				b.WriteString(help)
				if k+1 != len(cmd.Finch.Commands) {
					b.WriteString("\n")
				}
			}
		}
	} else {
		b.WriteString("Loaded commands:\n\n")

		for _, command := range cmd.Finch.Commands {
			help := command.Command.Help()

			if help.Description == "" {
				continue
			}

			b.WriteString(help.String(true))
		}
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, b.String())
	msg.ReplyToMessageID = update.Message.MessageID
	msg.ReplyMarkup = tgbotapi.ModeMarkdown
	return cmd.Finch.SendMessage(msg)
}
