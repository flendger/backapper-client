package command

import (
	"io"
	"log"
)

type Restart struct {
}

func (r *Restart) Run(request *Request) error {
	client := request.Client

	response, err := client.Get(Resolve(request))
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
