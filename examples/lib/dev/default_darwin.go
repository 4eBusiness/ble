package dev

import (
	"github.com/foureb/ble"
	"github.com/foureb/ble/darwin"
)

// DefaultDevice ...
func DefaultDevice() (d ble.Device, err error) {
	return darwin.NewDevice()
}
