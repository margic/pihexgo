package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
	"github.com/spf13/viper"
)

func main() {
	log.Info("Starting PiHexGo")

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

	// handle close
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		log.Info(sig)
		done <- true
	}()

	Init()
	defer Close()

	go run()

	// block until done
	<-done
	log.Debug("Shutting down.")
}

func run() {
	pin, err := embd.NewDigitalPin(18)

	if err != nil {
		log.Fatal("Error getting pin", err)
	}

	pin.SetDirection(embd.Out)

	ticker := time.NewTicker(time.Second * 1)
	off := true
	for range ticker.C {
		if off {
			err = pin.Write(embd.High)
			off = false
			log.Debug("On")
		} else {
			err = pin.Write(embd.Low)
			off = true
			log.Debug("Off")
		}
		if err != nil {
			log.Fatal("Error writing to pin", err)
		}
	}
}

// Init initialize the raspberry pi
func Init() {
	log.Info("Pi Time, initializing the Pi")
	embd.InitGPIO()
}

// Close clean up resources on shutdown
func Close() {
	log.Info("Cleaning up")
	defer embd.CloseGPIO()
}
