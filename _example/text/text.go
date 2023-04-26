package main

import (
	"errors"
	"os"

	"github.com/furzoom/log"
	"github.com/furzoom/log/handlers/text"
	e "github.com/pkg/errors"
)

func main() {
	log.SetHandler(text.New(os.Stderr))
	log.SetLevel(log.DebugLevel)
	ctx := log.WithFields(log.Fields{
		"file": "something.png",
		"type": "image/png",
		"user": "furzoom",
	})

	ctx.Debug("prepare figure")
	ctx.Info("upload")
	ctx.Info("upload complete")
	ctx.Warn("upload retry")
	ctx.WithError(errors.New("unauthorized")).Error("upload failed")
	err := errors.New("unknown error")
	ctx.WithError(e.Wrap(err, "test error")).Error("upload failed")
	ctx.Errorf("failed to upload %s", "img.png")
}
