package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hybridgroup/gobot"
)

func main() {
	log.Info("Starting PiHexGo")

	gBot := gobot.NewGobot()
	if gBot == nil {
		log.Fatal("Unable to get Gobot")
	}
}
