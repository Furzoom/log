package main

import (
	"errors"

	"github.com/furzoom/log"
)

func main() {
	ctx := log.WithFields(log.Fields{
		"file": "something.png",
		"type": "image/png",
		"user": "Furzoom",
	})

	ctx.Info("upload")
	ctx.Info("upload complete")
	ctx.Warn("upload retry")
	ctx.WithError(errors.New("unauthorized")).Error("upload failed")
	ctx.Errorf("failed to upload %s", "img.png")
}
