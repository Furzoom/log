package main

import (
	"errors"
	"os"
	"time"

	"github.com/furzoom/log"
	"github.com/furzoom/log/handlers/json"
)

func main() {
	log.SetHandler(json.New(os.Stderr))
	ctx := log.WithFields(log.Fields{
		"file": "something.png",
		"type": "image/png",
		"user": "furzoom",
	})

	for range time.Tick(time.Millisecond * 200) {
		ctx.Info("upload")
		ctx.Info("upload complete")
		ctx.Warn("upload retry")
		ctx.WithError(errors.New("unauthorized")).Error("upload failed")
	}
}