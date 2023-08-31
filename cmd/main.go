package main

import (
	"fmt"
	"strings"

	"github.com/sstallion/go-hid"
)

func fmtRelease(n uint16) string {
	if n == 0 {
		return "(empty)"
	}
	return fmt.Sprintf("%#04x (%x.%x)", n, n>>8, n&0xff)
}

func fmtString(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return "(empty)"
	}
	return s
}

func main() {
	hid.Enumerate(0, 0, func(info *hid.DeviceInfo) error {
		fmt.Printf("%s: ID %04x:%04x %s %s\n",
			info.Path, info.VendorID, info.ProductID, info.MfrStr, info.ProductStr)
		fmt.Println("Device Information:")
		fmt.Printf("\tPath         %s\n", info.Path)
		fmt.Printf("\tVendorID     %#04x\n", info.VendorID)
		fmt.Printf("\tProductID    %#04x\n", info.ProductID)
		fmt.Printf("\tSerialNbr    %s\n", fmtString(info.SerialNbr))
		fmt.Printf("\tReleaseNbr   %s\n", fmtRelease(info.ReleaseNbr))
		fmt.Printf("\tMfrStr       %s\n", fmtString(info.MfrStr))
		fmt.Printf("\tProductStr   %s\n", fmtString(info.ProductStr))
		fmt.Printf("\tUsagePage    %#04x\n", info.UsagePage)
		fmt.Printf("\tUsage        %#04x\n", info.Usage)
		fmt.Printf("\tInterfaceNbr %d\n", info.InterfaceNbr)
		fmt.Printf("\tBusType      %s\n", info.BusType)
		fmt.Println()
		return nil
	})
}
