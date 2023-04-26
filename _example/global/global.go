package main

import (
	"errors"
	"os"

	"github.com/furzoom/log"
	"github.com/furzoom/log/handlers/text"
)

func main() {
	log.SetHandler(text.New(os.Stderr))
	log.SetLevel(log.DebugLevel)
	log.SetDepth(2)
	log.Debug("prepare figure")
	log.Info("upload")
	log.Info("upload complete")
	log.Warn("upload retry")
	log.
		WithError(errors.New("unauthorized")).
		Error("upload failed")
	log.Errorf("failed to upload %s", "img.png")
}
