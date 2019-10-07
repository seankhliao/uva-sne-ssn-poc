package main

import (
	"log"

	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/bluez/profile/adapter"
	"github.com/muka/go-bluetooth/bluez/profile/device"
)

func main() {
	defer api.Exit()

	a, err := adapter.GetDefaultAdapter()
	if err != nil {
		log.Fatal(err)
	}

	a.FlushDevices()

	discovery, cancel, err := api.Discover(a, nil)
	if err != nil {
		log.Fatal("discover: ", err)
	}
	defer cancel()

	go func() {
		for ev := range discovery {
			if ev.Type == adapter.DeviceRemoved {
				continue
			}

			dev, err := device.NewDevice1(ev.Path)
			if err != nil {
				log.Println("newdevice: ", err)
				continue
			}
			if dev == nil {
				log.Println("dev nil", err)
				continue
			}

		}
	}()
	select {}
}
