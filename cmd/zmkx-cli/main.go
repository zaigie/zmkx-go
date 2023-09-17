package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zaigie/zmkx-go/zmkx"
)

func display(result any, err error) {
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	json, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println(string(json))
}
func main() {
	var rootCmd = &cobra.Command{Use: "zmkx-cli"}
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	var cmdVersion = &cobra.Command{
		Use:   "version",
		Short: "Print the version of devices",
		Run: func(cmd *cobra.Command, args []string) {
			devices := zmkx.FindDevices()
			if len(devices) == 0 {
				fmt.Println("No devices found")
				os.Exit(0)
			}
			for _, device := range devices {
				fmt.Println("=================================")
				fmt.Println("Device:", device.Name)
				display(device.GetVersion())
				fmt.Println("=================================")
			}
		},
	}
	var cmdKnob = &cobra.Command{
		Use:   "knob",
		Short: "Print the knob config",
		Run: func(cmd *cobra.Command, args []string) {
			devices := zmkx.FindDevices()
			if len(devices) == 0 {
				fmt.Println("No devices found")
				os.Exit(0)
			}
			for _, device := range devices {
				fmt.Println("=================================")
				fmt.Println("Device:", device.Name)
				display(device.GetKnobConfig())
				fmt.Println("=================================")
			}
		},
	}
	var cmdMotor = &cobra.Command{
		Use:   "motor",
		Short: "Print the motor state",
		Run: func(cmd *cobra.Command, args []string) {
			devices := zmkx.FindDevices()
			if len(devices) == 0 {
				fmt.Println("No devices found")
				os.Exit(0)
			}
			for _, device := range devices {
				fmt.Println("=================================")
				fmt.Println("Device:", device.Name)
				display(device.GetMotorState())
				fmt.Println("=================================")
			}
		},
	}
	var cmdRgb = &cobra.Command{
		Use:   "rgb",
		Short: "Print the RGB state",
		Run: func(cmd *cobra.Command, args []string) {
			devices := zmkx.FindDevices()
			if len(devices) == 0 {
				fmt.Println("No devices found")
				os.Exit(0)
			}
			for _, device := range devices {
				fmt.Println("=================================")
				fmt.Println("Device:", device.Name)
				display(device.GetRgbState())
				fmt.Println("=================================")
			}
		},
	}
	var cmdEink = &cobra.Command{
		Use:     "eink",
		Short:   "Set the Eink image",
		Example: `zmkx-cli eink -f /Users/zaigie/Desktop/keyboard_ink.jpg`,
		Run: func(cmd *cobra.Command, args []string) {
			devices := zmkx.FindDevices()
			if len(devices) == 0 {
				fmt.Println("No devices found")
				os.Exit(0)
			}
			file, _ := cmd.Flags().GetString("file")
			threshold, _ := cmd.Flags().GetInt("threshold")
			if threshold < 0 || threshold > 65535 {
				fmt.Println("Threshold must be between 0 and 65535.")
				os.Exit(1)
			}
			for _, device := range devices {
				fmt.Println("=================================")
				fmt.Println("Device:", device.Name)
				imageBytes, err := zmkx.LoadImage(file, uint16(threshold))
				if err != nil {
					fmt.Println(err)
					return
				}
				display(device.SetEinkImage(imageBytes))
				fmt.Println("=================================")
			}
		},
	}
	cmdEink.Flags().StringP("file", "f", "", "Eink image filename")
	cmdEink.Flags().IntP("threshold", "t", 0, "Eink image threshold (1-65535) (default 32768)")
	rootCmd.AddCommand(cmdVersion, cmdKnob, cmdMotor, cmdRgb, cmdEink)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
