package main

import (
	"cpu-temp-Polybar/src/helpers"
	"cpu-temp-Polybar/src/parsers"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	// Assign Polybar Color from Args
	args := helpers.ParseInput()
	iconClr := args.IconColor
	strClr := args.StrColor
	// cpuType := args.CpuType

	tempEmojis := []string{"", "", "", "", ""}
	var tempInfo parsers.TempInfo

	for {
		// RUN COMMAND TO GET INFO
		cmd := exec.Command("sensors", "-u")
		out, err := cmd.Output()

		if err != nil {
			fmt.Println("Command Error:", err.Error())
			os.Exit(1)
		}

		// PARSE DATA FROM INFO
		parsers.ParseOutput_intel(out, &tempInfo)

		// CHOOSE APPROPRIATE EMOJI
		tempStatusE := tempEmojis[0] // Temp < 40C
		switch val := tempInfo.PackageTemp; {
		case val > 40 && val < 65:
			tempStatusE = tempEmojis[1]
		case val < 70:
			tempStatusE = tempEmojis[2]
		case val <= 90:
			tempStatusE = tempEmojis[3]
		case val > 90:
			tempStatusE = tempEmojis[4]
		}

		fmt.Printf("%%{F%s}%s %%{F%s}%.1f°C\n", iconClr, tempStatusE, strClr, tempInfo.PackageTemp)
		time.Sleep(5 * time.Second)
	}
}
