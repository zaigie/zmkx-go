package zmkx

import (
	"github.com/sstallion/go-hid"
	"github.com/zaigie/zmkx-go/internal/img"
	u "github.com/zaigie/zmkx-go/internal/utils"
)

func filterDevices(devices []*ZMKXDevice, features []string) []*ZMKXDevice {
	if len(features) == 0 {
		return devices
	}
	var filtered []*ZMKXDevice
	for _, device := range devices {
		deviceVersion, err := device.GetVersion()
		if err != nil {
			break
		}
		deviceFeatures := deviceVersion["features"].(map[string]interface{})
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
	hid.Enumerate(u.ZmkxVID, u.ZmkxPID, func(info *hid.DeviceInfo) error {
		if info.UsagePage == u.ZmkxUsage {
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
	return img.LoadImage(filename, threshold)
}
