package main

import (
	"fmt"

	"github.com/zaigie/zmkx-go/zmkx"
)

func main() {
	devices := zmkx.FindDevices()
	fmt.Println(devices)
}
