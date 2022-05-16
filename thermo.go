package main

import (
	"github.com/MichaelS11/go-dht"
)

type Thermo struct {
	Temperature float64
	Humidity    float64
}

func readData() (*Thermo, error) {
	err := dht.HostInit()
	if err != nil {
		return nil, err
	}

	dht, err := dht.NewDHT("GPIO4", dht.Celsius, "")
	if err != nil {
		return nil, err
	}

	humidity, temperature, err := dht.ReadRetry(10)
	if err != nil {
		return nil, err
	}

	status := Thermo{temperature, humidity}

	return &status, nil
}
