package main

import (
	"backapper-client/command"
	"backapper-client/params"
	"backapper-client/resolver"
	"log"
)

func main() {
	//command appName filePath
	commandParams, err := params.Resolve()
	if err != nil {
		log.Println(err.Error())
		return
	}

	commandResolver := resolver.New(command.Backup{})
	resolvedCommand, err := commandResolver.GetCommand(commandParams)
	if err != nil {
		log.Println(err)
		return
	}

	err = resolvedCommand.Run(commandParams)
	if err != nil {
		log.Println(err)
	}
}
