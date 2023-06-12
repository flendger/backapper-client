package main

import (
	"backapper-client/clientfactory"
	"backapper-client/command"
	"backapper-client/params"
	"backapper-client/profileholder"
	"backapper-client/resolver"
	"log"
)

const configPath = "backapper-client.json"
const version = "2.0.0"

var logger *log.Logger

func init() {
	logger = log.Default()
}

func main() {
	logger.Println("version:", version)

	holder := profileholder.Read(configPath, logger)

	//-c=command -p=profile
	commandParams, err := params.Resolve()
	if err != nil {
		log.Println(err.Error())
		return
	}

	profile, err := holder.GetProfile(commandParams.Profile)
	if err != nil {
		log.Println(err)
		return
	}

	client := clientfactory.CreateClient(profile)

	request := command.NewRequest(commandParams, profile, client)

	commandResolver := resolver.New(&command.Backup{}, &command.Deploy{}, &command.Restart{})
	resolvedCommand, err := commandResolver.GetCommand(commandParams)
	if err != nil {
		log.Println(err)
		return
	}

	err = resolvedCommand.Run(request)
	if err != nil {
		log.Println(err)
	}
}
