package finch

import (
	"bytes"
	"github.com/syfaro/telegram-bot-api"
)

// Help contains information about a command,
// used for showing info in the help command.
type Help struct {
	Name        string
	Description string
	Example     string
}

// String converts a Help struct into a pretty string.
//
// full makes each command item multiline with extra newlines.
func (h Help) String(full bool) string {
	b := &bytes.Buffer{}

	b.WriteString(h.Name)
	if full {
		b.WriteString("\n")
	} else {
		b.WriteString(" - ")
	}
	b.WriteString(h.Description)
	b.WriteString("\n")

	if full {
		b.WriteString("Example: ")
		b.WriteString(h.Example)
		b.WriteString("\n")
	}

	b.WriteString("\n")

	return b.String()
}

// Command contains the methods a must have.
type Command interface {
	Help() Help
	Init(*CommandState, *Finch) error
	ShouldExecute(tgbotapi.Update) bool
	Execute(tgbotapi.Update) error
	ExecuteKeyboard(tgbotapi.Update) error
}

// CommandBase is a default Command that handles various tasks for you,
// and allows for you to not have to write empty methods.
type CommandBase struct {
	MyState *CommandState
	Finch   *Finch
}

// Help returns an empty Help struct.
func (CommandBase) Help() Help { return Help{} }

// Init sets MyState equal to the current CommandState.
//
// If you overwrite this method, you should still set MyState equal to CommandState!
func (cmd *CommandBase) Init(c *CommandState, f *Finch) error {
	cmd.MyState = c
	cmd.Finch = f

	return nil
}

// ShouldExecute returns false, you should overwrite this method.
func (CommandBase) ShouldExecute(tgbotapi.Update) bool { return false }

// Execute returns nil to show no error, you should overwrite this method.
func (CommandBase) Execute(tgbotapi.Update) error { return nil }

// ExecuteKeyboard returns nil to show no error, you may overwrite this when
// you are expecting to get a reply that is not a command.
func (CommandBase) ExecuteKeyboard(tgbotapi.Update) error { return nil }

// Get fetches an item from the Config struct.
func (cmd CommandBase) Get(key string) interface{} {
	return cmd.Finch.Config[key]
}

// Set sets an item in the Config struct, then saves it.
func (cmd CommandBase) Set(key string, value interface{}) {
	cmd.Finch.Config[key] = value
	cmd.Finch.Config.Save()
}

// CommandState is the current state of a command.
// It contains the command and if the command is waiting for a reply.
type CommandState struct {
	Command         Command
	WaitingForReply bool
}