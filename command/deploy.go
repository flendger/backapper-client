package command

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"os"
)

type Deploy struct {
}

func (d *Deploy) Run(r *Request) error {
	profile := r.Profile
	path := profile.Path
	client := r.Client

	log.Println("Starting deploy: " + profile.App)
	log.Println("Opening file: " + path)
	src, err := os.Open(path)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	file, err := w.CreateFormFile("file", path)
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

	log.Println("Uploading file...")
	response, err := client.Post(Resolve(r), w.FormDataContentType(), &buf)
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
