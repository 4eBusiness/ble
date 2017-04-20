package dev

import (
	"github.com/foureb/ble"
	"github.com/foureb/ble/linux"
)

// DefaultDevice ...
func DefaultDevice() (d ble.Device, err error) {
	return linux.NewDevice()
}
