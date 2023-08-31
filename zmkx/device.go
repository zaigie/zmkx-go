package zmkx

import "github.com/sstallion/go-hid"

type ZMKXDevice struct {
	hid.Device
	Path      string
	Usage     uint16
	SerialNbr string
}
