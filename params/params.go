package params

import (
	"errors"
	"os"
)

type Params struct {
	Address  string
	Command  string
	AppName  string
	FilePath string
}

func Resolve() (*Params, error) {
	if len(os.Args) < 2 {
		return nil, errors.New("usage: address command app file (optional)")
	}
	address := os.Args[1]

	if len(os.Args) < 3 {
		return nil, errors.New("no command passed")
	}
	command := os.Args[2]

	if len(os.Args) < 4 {
		return nil, errors.New("no app name passed")
	}
	appName := os.Args[3]

	var path string
	if len(os.Args) > 4 {
		path = os.Args[4]
	}

	return &Params{Address: address, Command: command, AppName: appName, FilePath: path}, nil
}
