package main

import (
	"os"

	"github.com/furzoom/log"
	"github.com/furzoom/log/handlers/logfmt"
	"github.com/pkg/errors"
)

func main() {
	log.SetHandler(logfmt.New(os.Stderr))

	filename := "something.png"
	body := []byte("whatever")

	ctx := log.WithField("filename", filename)
	err := upload(filename, body)
	if err != nil {
		ctx.WithError(err).Error("upload failed")
	}
}

func upload(name string, b []byte) error {
	err := put("/images/"+name, b)
	if err != nil {
		return errors.Wrap(err, "uploading to s3")
	}
	return nil
}

func put(key string, b []byte) error {
	return errors.New("unauthorized")
}
