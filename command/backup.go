package command

import (
	"io"
	"log"
)

type Backup struct {
}

func (b *Backup) Run(r *Request) error {
	client := r.Client

	response, err := client.Get(Resolve(r))
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
