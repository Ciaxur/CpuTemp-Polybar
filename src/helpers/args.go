package helpers

import "flag"

// Arguments Struct
type CliArguments struct {
	IconColor   string
	StrColor    string
	CpuType     string
	CpuTempType string
}

// Parses Argument Input with Default Values
func ParseInput() CliArguments {
	// COLOR OPTIONS
	var IconColor = flag.String("i", "#FF", "CPU Icon Color Hex Code")
	var StrColor = flag.String("s", "#FF", "CPU Temperature Color Hex Code")
	var CpuType = flag.String("c", "intel", "Cpu Type (intel/amd)")
	var CpuTempType = flag.String("t", "die", "Cpu Temperature type (die/tctl)")

	flag.Parse()
	return CliArguments{
		*IconColor,
		*StrColor,
		*CpuType,
		*CpuTempType,
	}
}
