package path

import (
	"errors"
	"fmt"
	"strings"
)

type CommandPath struct {
	CommandName string
	Electronics string
	Phones      string
}

var ErrUnknownCommand = errors.New("unknown command")

func ParseCommand(commandText string) (CommandPath, error) {
	commandParts := strings.SplitN(commandText, "_", 3)

	if len(commandParts) != 3 {
		return CommandPath{}, ErrUnknownCommand
	}

	return CommandPath{
		CommandName: commandParts[0],
		Electronics: commandParts[1],
		Phones:      commandParts[2],
	}, nil
}

func (c CommandPath) WithCommandName(commandName string) CommandPath {
	c.CommandName = commandName

	return c
}

func (c CommandPath) String() string {
	return fmt.Sprintf("/%s_%s_%s", c.CommandName, c.Electronics, c.Phones)
}
