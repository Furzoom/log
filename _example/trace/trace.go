package main

import (
	"github.com/furzoom/log"
	"github.com/furzoom/log/handlers/text"
	"os"
)

func work(ctx log.Interface) (err error) {
	path := "README.md"
	defer ctx.WithField("path", path).Trace("opening").Stop(&err)
	_, err = os.Open(path)
	return
}

func main() {
	log.SetHandler(text.New(os.Stderr))
	log.SetLevel(log.DebugLevel)

	ctx := log.WithFields(log.Fields{
		"app": "myapp",
		"evn": "product",
	})

	_ = work(ctx)
}
