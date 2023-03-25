package params

import (
	"errors"
	"os"
)

type Params struct {
	Command  string
	AppName  string
	FilePath string
}

func Resolve() (*Params, error) {
	if len(os.Args) < 2 {
		return nil, errors.New("no command passed")
	}
	command := os.Args[1]

	if len(os.Args) < 3 {
		return nil, errors.New("no app name passed")
	}
	appName := os.Args[2]

	var path string
	if len(os.Args) > 3 {
		path = os.Args[3]
	}

	return &Params{Command: command, AppName: appName, FilePath: path}, nil
}
