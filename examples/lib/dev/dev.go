package dev

import "github.com/foureb/ble"

// NewDevice ...
func NewDevice(impl string) (d ble.Device, err error) {
	return DefaultDevice()
}
