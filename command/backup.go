package command

import (
	"backapper-client/params"
	"backapper-client/pathresolver"
	"io"
	"log"
	"net/http"
)

type Backup struct {
}

func (b *Backup) Run(p *params.Params) error {
	response, err := http.Get(pathresolver.Resolve(p))
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(response.Body)

	resp, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	log.Println(string(resp))
	return nil
}

func (b *Backup) Type() string {
	return "backup"
}
