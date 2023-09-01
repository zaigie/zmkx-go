package zmkx

import (
	"github.com/sstallion/go-hid"
)

func filterDevices(devices []*ZMKXDevice, features []string) []*ZMKXDevice {
	if len(features) == 0 {
		return devices
	}
	var filtered []*ZMKXDevice
	for _, device := range devices {
		deviceFeatures := device.GetVersion()["features"].(map[string]interface{})
		for _, feature := range features {
			if val, ok := deviceFeatures[feature]; ok && val.(bool) {
				filtered = append(filtered, device)
				break
			}
		}
	}
	return filtered
}

func FindDevices(features ...string) []*ZMKXDevice {
	devices := make([]*ZMKXDevice, 0)
	hid.Enumerate(ZmkxVID, ZmkxPID, func(info *hid.DeviceInfo) error {
		if info.UsagePage == ZmkxUsage {
			devices = append(devices, &ZMKXDevice{
				Name:      info.MfrStr + " " + info.ProductStr,
				Path:      info.Path,
				Usage:     info.Usage,
				UsagePage: info.UsagePage,
				SerialNbr: info.SerialNbr,
			})
		}
		return nil
	})
	return filterDevices(devices, features)
}

func LoadImage(filename string, threshold uint16) ([]byte, error) {
	return loadImage(filename, threshold)
}
