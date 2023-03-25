package resolver

import (
	"backapper-client/command"
	"backapper-client/params"
	"errors"
)

type CommandResolver struct {
	commands map[string]command.Command
}

func (r *CommandResolver) GetCommand(params *params.Params) (command.Command, error) {
	c, exist := r.commands[params.Command]
	if !exist {
		return nil, errors.New("command doesn't exist: " + params.Command)
	}

	return c, nil
}

func New(commands ...command.Command) *CommandResolver {
	resolver := CommandResolver{
		commands: make(map[string]command.Command),
	}

	for _, curCommand := range commands {
		resolver.commands[curCommand.Type()] = curCommand
	}

	return &resolver
}
