package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/raspi"
)

func main() {
	log.Info("Starting PiHexGo")

	gbot := gobot.NewGobot()
	if gbot == nil {
		log.Fatal("Unable to get Gobot")
	}

	pi := raspi.NewRaspiAdaptor("raspi")
	pin := gpio.NewDirectPinDriver(pi, "pin", "18")

	work := func() {
		log.Info("Work")
		pin.DigitalWrite(1)
	}

	robot := gobot.NewRobot("pihexgo",
		[]gobot.Connection{pi},
		[]gobot.Device{pin},
		work,
	)
	gbot.AddRobot(robot)
	gbot.Start()
}
