package main

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/pihex/")
	viper.AddConfigPath("$HOME/.pihex")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		log.Fatal("Failed to read a config ", err)
	}

	debug := viper.GetBool("debug")
	if debug {
		log.SetLevel(log.DebugLevel)
	}

	log.Info("Starting PiHexGo")

	log.Debug("Init GPIO")

	embd.InitGPIO()
	defer embd.CloseGPIO()
	embd.SetDirection(18, embd.Out)

	ticker := time.NewTicker(time.Second * 1)
	off := true
	for range ticker.C {
		if off {
			embd.DigitalWrite(18, embd.High)
			off = false
			log.Debug("On")
		} else {
			embd.DigitalWrite(24, embd.Low)
			off = true
			log.Debug("Off")
		}
	}

}
