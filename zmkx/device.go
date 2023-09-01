package zmkx

import (
	"fmt"
	"log"

	"github.com/sstallion/go-hid"
	usbComm "github.com/zaigie/zmkx-go/comm"
	"google.golang.org/protobuf/proto"
)

type ZMKXDevice struct {
	Path      string
	Usage     uint16
	UsagePage uint16
	SerialNbr string
}

func (d *ZMKXDevice) call(h2d *usbComm.MessageH2D) *usbComm.MessageD2H {
	msgOut, err := DelimitedEncode(h2d)
	if err != nil {
		log.Fatalf("Failed to encode h2d message: %v", err)
	}

	opendDevice, err := hid.OpenPath(d.Path)
	if err != nil {
		log.Fatalf("Failed to open device: %v", err)
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
		n, err := opendDevice.Read(readBuffer)
		if err != nil {
			log.Fatalf("Failed to read from device: %v", err)
		}
		fmt.Printf("Read %d bytes\n", n)

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
		log.Fatalf("Failed to decode d2h message: %v", err)
	}
	d2h := &usbComm.MessageD2H{}
	if err := proto.Unmarshal(received, d2h); err != nil {
		log.Fatalf("Failed to unmarshal d2h message: %v", err)
	}

	return d2h
}

func (d *ZMKXDevice) GetVersion() map[string]interface{} {
	action := usbComm.Action_VERSION
	h2d := &usbComm.MessageH2D{
		Action: &action,
	}
	d2h := d.call(h2d)
	return proto2Map(d2h.GetVersion())
}

// Knob
func (d *ZMKXDevice) GetKnobConfig() map[string]interface{} {
	action := usbComm.Action_KNOB_GET_CONFIG
	h2d := &usbComm.MessageH2D{
		Action: &action,
	}
	d2h := d.call(h2d)
	return proto2Map(d2h.GetKnobConfig())
}

// func (d *ZMKXDevice) SetKnobConfig(config *usbComm.KnobConfig) map[string]interface{} {
// 	action := usbComm.Action_KNOB_SET_CONFIG
// 	h2d := &usbComm.MessageH2D{
// 		Action: &action,
// 		Payload: &usbComm.MessageH2D_KnobConfig{
// 			KnobConfig: config,
// 		},
// 	}
// 	d2h := d.call(h2d)
// 	return proto2Map(d2h.GetKnobConfig())
// }

// Motor
func (d *ZMKXDevice) GetMotorState() map[string]interface{} {
	action := usbComm.Action_MOTOR_GET_STATE
	h2d := &usbComm.MessageH2D{
		Action: &action,
	}
	d2h := d.call(h2d)
	return proto2Map(d2h.GetMotorState())
}

// RGB
func (d *ZMKXDevice) GetRgbState() map[string]interface{} {
	action := usbComm.Action_RGB_GET_STATE
	h2d := &usbComm.MessageH2D{
		Action: &action,
	}
	d2h := d.call(h2d)
	return proto2Map(d2h.GetRgbState())
}

// Eink
func (d *ZMKXDevice) SetEinkImage(imageBytes []byte) map[string]interface{} {
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
	d2h := d.call(h2d)
	return proto2Map(d2h.GetEinkImage())
}
