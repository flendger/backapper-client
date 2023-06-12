package params

import (
	"errors"
	"flag"
)

type Params struct {
	Command string
	Profile string
}

func Resolve() (*Params, error) {

	var commandName string
	flag.StringVar(&commandName, "c", "", "command")

	var profileName string
	flag.StringVar(&profileName, "p", "", "profile")

	flag.Parse()

	if commandName == "" {
		return nil, errors.New("command missing")
	}

	if profileName == "" {
		return nil, errors.New("profile missing")
	}

	return &Params{Command: commandName, Profile: profileName}, nil
}
