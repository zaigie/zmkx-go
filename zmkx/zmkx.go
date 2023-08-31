package zmkx

import (
	"fmt"

	"github.com/sstallion/go-hid"
)

func Hex2String(n uint16) string {
	return fmt.Sprintf("%#04x", n)
}

func FindDevices() []*ZMKXDevice {
	devices := make([]*ZMKXDevice, 0)
	hid.Enumerate(0, 0, func(info *hid.DeviceInfo) error {
		if fmt.Sprintf("%#04x", info.ProductID) == ZMKX_PID && fmt.Sprintf("%#04x", info.VendorID) == ZMKX_VID && fmt.Sprintf("%#04x", info.UsagePage) == ZMKX_USAGE {
			devices = append(devices, &ZMKXDevice{
				Path:      info.Path,
				Usage:     info.Usage,
				SerialNbr: info.SerialNbr,
			})
		}
		return nil
	})
	return devices
}
