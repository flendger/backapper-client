package profileholder

import (
	"backapper-client/profile"
	"encoding/json"
	"log"
	"os"
)

func Read(configFile string, logger *log.Logger) *Holder {

	bytes, err := os.ReadFile(configFile)
	if err != nil {
		logger.Println("Couldn't read profile file:", err)
		return nil
	}

	var profiles []*profile.AppProfile
	err = json.Unmarshal(bytes, &profiles)
	if err != nil {
		logger.Println("Couldn't parse profiles", configFile, err)
		return NewHolder()
	}

	return NewHolder(profiles...)
}
