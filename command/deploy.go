package command

import (
	"backapper-client/params"
	"backapper-client/pathresolver"
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

type Deploy struct {
}

func (d *Deploy) Run(p *params.Params) error {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	file, err := w.CreateFormFile("file", p.FilePath)
	if err != nil {
		return err
	}

	src, err := os.Open(p.FilePath)
	if err != nil {
		return err
	}

	_, errCopy := io.Copy(file, src)
	if errCopy != nil {
		return errCopy
	}

	errClose := w.Close()
	if errClose != nil {
		return errClose
	}

	response, err := http.Post(pathresolver.Resolve(p), w.FormDataContentType(), &buf)
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

func (d *Deploy) Type() string {
	return "deploy"
}
