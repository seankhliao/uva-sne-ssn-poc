package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/paypal/gatt"
	"github.com/paypal/gatt/examples/option"
)

func init() {
	// log format, controlled by LOGFMT
	logfmt := os.Getenv("LOGFMT")
	if logfmt != "json" {
		logfmt = "text"
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: !terminal.IsTerminal(int(os.Stdout.Fd()))})
	}
	log.Info().Str("FMT", logfmt).Msg("log format")

	// log level, controlled by LOGLVL
	level, err := zerolog.ParseLevel(os.Getenv("LOGLVL"))
	if err != nil || level == zerolog.NoLevel {
		level = zerolog.InfoLevel
	}
	log.Info().Str("LOGLVL", level.String()).Msg("log level")
	zerolog.SetGlobalLevel(level)
}

func main() {
	//
	d, err := gatt.NewDevice(option.DefaultClientOptions...)
	if err != nil {
		log.Fatal().Err(err).Msg("gatt.NewDevice")
	}

	// Register handlers.
	d.Handle(gatt.PeripheralDiscovered(onPeriphDiscovered))
	d.Init(onStateChanged)
	select {}
}
func onStateChanged(d gatt.Device, s gatt.State) {
	fmt.Println("State:", s)
	switch s {
	case gatt.StatePoweredOn:
		fmt.Println("scanning...")
		d.Scan([]gatt.UUID{}, false)
		return
	default:
		d.StopScanning()
	}
}

func onPeriphDiscovered(p gatt.Peripheral, a *gatt.Advertisement, rssi int) {
	// if p.ID() == "D8:68:C3:8C:78:CE" || p.ID() == "c0:ee:fb:ff:33:cf" {
	fmt.Printf("\nPeripheral ID:%s, NAME:(%s)\n", p.ID(), p.Name())
	fmt.Println("  Local Name        =", a.LocalName)
	fmt.Println("  TX Power Level    =", a.TxPowerLevel)
	fmt.Println("  Manufacturer Data =", a.ManufacturerData)
	fmt.Println("  Service Data      =", a.ServiceData)
	// }
}
