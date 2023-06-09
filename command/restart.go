package command

import (
	"backapper-client/params"
	"backapper-client/pathresolver"
	"io"
	"log"
	"net/http"
)

type Restart struct {
}

func (r *Restart) Run(p *params.Params) error {
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

func (r *Restart) Type() string {
	return "restart"
}
