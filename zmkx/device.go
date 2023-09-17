package zmkx

import (
	"errors"

	"github.com/sstallion/go-hid"
	usbComm "github.com/zaigie/zmkx-go/comm"
	"google.golang.org/protobuf/proto"
)

type ZMKXDevice struct {
	Name      string
	Path      string
	Usage     uint16
	UsagePage uint16
	SerialNbr string
}

func (d *ZMKXDevice) call(h2d *usbComm.MessageH2D) (*usbComm.MessageD2H, error) {
	msgOut, err := DelimitedEncode(h2d)
	if err != nil {
		return nil, errors.New("Failed to encode h2d message: " + err.Error())
	}

	opendDevice, err := hid.OpenPath(d.Path)
	if err != nil {
		return nil, errors.New("Failed to open device: " + err.Error())
	}
	defer opendDevice.Close()

	for offset := 0; offset < len(msgOut); offset += PayloadSize {
		buf := msgOut[offset:minInt(offset+PayloadSize, len(msgOut))]
		hdr := []byte{byte(d.Usage), byte(len(buf))}
		payload := append(hdr, Ljust(buf, 0, PayloadSize)...)
		opendDevice.Write(payload)
	}

	msgIn := make([]byte, 0)
	readBuffer := make([]byte, 1+ReportCount)

	for {
		_, err := opendDevice.Read(readBuffer)
		if err != nil {
			return nil, errors.New("Failed to read from device: " + err.Error())
		}
		// fmt.Printf("Read %d bytes\n", n)

		// Assuming the device supports multiple reports, so we skip the first byte (report ID)
		cnt := readBuffer[1]
		msgIn = append(msgIn, readBuffer[2:cnt+2]...) // append read bytes to msgIn
		if cnt < byte(PayloadSize) {
			break
		}
	}

	// Decode the received message
	received, err := DelimitedDecode(msgIn)
	if err != nil {
		return nil, errors.New("Failed to decode d2h message: " + err.Error())
	}
	d2h := &usbComm.MessageD2H{}
	if err := proto.Unmarshal(received, d2h); err != nil {
		return nil, errors.New("Failed to unmarshal d2h message: " + err.Error())
	}

	return d2h, nil
}

func (d *ZMKXDevice) GetVersion() (map[string]interface{}, error) {
	action := usbComm.Action_VERSION
	h2d := &usbComm.MessageH2D{
		Action: &action,
	}
	d2h, err := d.call(h2d)
	if err != nil {
		return nil, err
	}
	result, err := proto2Map(d2h.GetVersion())
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Knob
func (d *ZMKXDevice) GetKnobConfig() (map[string]interface{}, error) {
	action := usbComm.Action_KNOB_GET_CONFIG
	h2d := &usbComm.MessageH2D{
		Action: &action,
	}
	d2h, err := d.call(h2d)
	if err != nil {
		return nil, err
	}
	result, err := proto2Map(d2h.GetKnobConfig())
	if err != nil {
		return nil, err
	}
	return result, nil
}

// func (d *ZMKXDevice) SetKnobConfig(config *usbComm.KnobConfig) (map[string]interface{}, error) {
// 	action := usbComm.Action_KNOB_SET_CONFIG
// 	h2d := &usbComm.MessageH2D{
// 		Action: &action,
// 		Payload: &usbComm.MessageH2D_KnobConfig{
// 			KnobConfig: config,
// 		},
// 	}
// 	d2h, err := d.call(h2d)
// 	if err != nil {
// 		return nil, err
// 	}
// 	result, err := proto2Map(d2h.GetKnobConfig())
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

// Motor
func (d *ZMKXDevice) GetMotorState() (map[string]interface{}, error) {
	action := usbComm.Action_MOTOR_GET_STATE
	h2d := &usbComm.MessageH2D{
		Action: &action,
	}
	d2h, err := d.call(h2d)
	if err != nil {
		return nil, err
	}
	result, err := proto2Map(d2h.GetMotorState())
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RGB
func (d *ZMKXDevice) GetRgbState() (map[string]interface{}, error) {
	action := usbComm.Action_RGB_GET_STATE
	h2d := &usbComm.MessageH2D{
		Action: &action,
	}
	d2h, err := d.call(h2d)
	if err != nil {
		return nil, err
	}
	result, err := proto2Map(d2h.GetRgbState())
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Eink
func (d *ZMKXDevice) SetEinkImage(imageBytes []byte) (map[string]interface{}, error) {
	action := usbComm.Action_EINK_SET_IMAGE
	image := &usbComm.EinkImage{
		Id:   genImageID(),
		Bits: imageBytes,
	}
	h2d := &usbComm.MessageH2D{
		Action: &action,
		Payload: &usbComm.MessageH2D_EinkImage{
			EinkImage: image,
		},
	}
	d2h, err := d.call(h2d)
	if err != nil {
		return nil, err
	}
	result, err := proto2Map(d2h.GetEinkImage())
	if err != nil {
		return nil, err
	}
	return result, nil
}
